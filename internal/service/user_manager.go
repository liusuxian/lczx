package service

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/utility/utils"
)

type sUserManager struct {
	notCheckAuthUserIds *gset.Set // 无需验证权限的用户ID
}

var (
	userManagerCtx      = gctx.New()
	notCheckAuthUserIds = g.Cfg().MustGet(userManagerCtx, "userManager.notCheckAuthUserIds").Interfaces()
	insUserManager      = sUserManager{
		notCheckAuthUserIds: gset.NewFrom(notCheckAuthUserIds),
	}
)

// UserManager 用户管理服务
func UserManager() *sUserManager {
	return &insUserManager
}

// NotCheckAuthUserIds 获取无需验证权限的用户ID
func (s *sUserManager) NotCheckAuthUserIds() *gset.Set {
	return s.notCheckAuthUserIds
}

// GetUserList 获取用户列表
func (s *sUserManager) GetUserList(ctx context.Context, req *v1.UserListReq) (total int, list []*entity.User, err error) {
	deptIdsMap := gmap.New()
	if req.DeptId != "" {
		// 获取部门状态为正常的部门列表
		var depts []*entity.Dept
		depts, err = Dept().GetStatusEnableDepts(ctx)
		if err != nil {
			return
		}
		deptId := gconv.Uint64(req.DeptId)
		deptIdsMap.Set(deptId, true)
		Dept().FindSonIdsByParentId(depts, deptId, deptIdsMap)
	}
	model := dao.User.Ctx(ctx)
	columns := dao.User.Columns()
	order := "id ASC"
	if !deptIdsMap.IsEmpty() {
		model = model.WhereIn(columns.DeptId, deptIdsMap.Keys())
	}
	if req.Passport != "" {
		model = model.WhereLike(columns.Passport, "%"+req.Passport+"%")
	}
	if req.Realname != "" {
		model = model.WhereLike(columns.Realname, "%"+req.Realname+"%")
	}
	if req.Mobile != "" {
		model = model.WhereLike(columns.Mobile, "%"+req.Mobile+"%")
	}
	if req.Status != "" {
		model = model.Where(columns.Status, req.Status)
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.CreatedAt, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.CreatedAt, req.EndTime)
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).Order(order).Scan(&list)
	return
}

// AddUser 新增用户
func (s *sUserManager) AddUser(ctx context.Context, req *v1.UserAddReq) (err error) {
	// 生成随机密码盐
	salt := grand.S(10)
	// 密码加密
	password := utils.EncryptPassword(req.Password, salt)
	// 开启事务
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查用户账号是否可用
		var available bool
		var terr error
		available, terr = s.IsPassportAvailable(ctx, req.Passport)
		if terr != nil {
			return terr
		}
		if !available {
			return gerror.Newf(`账号[%s]已存在`, req.Passport)
		}
		// 检查用户手机是否可用
		available, terr = s.IsMobileAvailable(ctx, req.Mobile)
		if terr != nil {
			return terr
		}
		if !available {
			return gerror.Newf(`手机[%s]已存在`, req.Mobile)
		}
		// 通过部门ID判断部门信息是否存在
		var deptExists bool
		deptExists, terr = Dept().DeptExistsById(ctx, req.DeptId)
		if terr != nil {
			return terr
		}
		if !deptExists {
			return gerror.Newf(`部门ID[%d]不存在`, req.DeptId)
		}
		// 保存用户数据
		var userId int64
		userId, terr = s.saveUser(ctx, req, password, salt)
		if terr != nil {
			return terr
		}
		// 通过角色ID列表获取可用角色
		var enableRoles []*entity.Role
		enableRoles, terr = Role().GetEnableRolesByIds(ctx, req.RoleIds)
		if terr != nil {
			return terr
		}
		if len(enableRoles) == 0 {
			return gerror.Newf(`角色ID列表%v不存在`, req.RoleIds)
		}
		roleIdsMap := gmap.New()
		for _, r := range enableRoles {
			roleIdsMap.Set(r.Id, true)
		}
		// 处理客户端发送过来的角色ID列表
		for _, rid := range req.RoleIds {
			if !roleIdsMap.Contains(rid) {
				return gerror.Newf(`角色ID[%d]不存在`, rid)
			}
		}
		// 添加用户角色信息
		terr = Role().AddUserRoles(ctx, req.RoleIds, gconv.Uint64(userId))
		return terr
	})
	return
}

// EditUser 编辑用户
func (s *sUserManager) EditUser(ctx context.Context, req *v1.UserEditReq) (err error) {
	// 开启事务
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查用户手机是否可用
		var user *entity.User
		var terr error
		user, terr = User().GetUserById(ctx, req.Id)
		if terr != nil {
			return terr
		}
		if user == nil {
			return gerror.Newf(`用户ID[%d]不存在`, req.Id)
		}
		if user.Mobile != req.Mobile {
			var available bool
			available, terr = s.IsMobileAvailable(ctx, req.Mobile)
			if terr != nil {
				return terr
			}
			if !available {
				return gerror.Newf(`手机[%s]已存在`, req.Mobile)
			}
		}
		// 通过部门ID判断部门信息是否存在
		var deptExists bool
		deptExists, terr = Dept().DeptExistsById(ctx, req.DeptId)
		if terr != nil {
			return terr
		}
		if !deptExists {
			return gerror.Newf(`部门ID[%d]不存在`, req.DeptId)
		}
		// 更新用户数据
		terr = s.updateUser(ctx, req)
		if terr != nil {
			return terr
		}
		// 通过角色ID列表获取可用角色
		var enableRoles []*entity.Role
		enableRoles, terr = Role().GetEnableRolesByIds(ctx, req.RoleIds)
		if terr != nil {
			return terr
		}
		if len(enableRoles) == 0 {
			return gerror.Newf(`角色ID列表%v不存在`, req.RoleIds)
		}
		roleIdsMap := gmap.New()
		for _, r := range enableRoles {
			roleIdsMap.Set(r.Id, true)
		}
		// 处理客户端发送过来的角色ID列表
		for _, rid := range req.RoleIds {
			if !roleIdsMap.Contains(rid) {
				return gerror.Newf(`角色ID[%d]不存在`, rid)
			}
		}
		// 修改用户角色信息
		terr = Role().EditUserRoles(ctx, req.RoleIds, req.Id)
		return terr
	})
	return
}

// ResetUserPwd 重置用户密码
func (s *sUserManager) ResetUserPwd(ctx context.Context, id uint64, newPassword string) (err error) {
	// 生成随机密码盐
	salt := grand.S(10)
	newEncryptPassword := utils.EncryptPassword(newPassword, salt)
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Salt:     salt,
		Password: newEncryptPassword,
	}).Where(do.User{Id: id}).Update()
	return
}

// SetUserStatus 设置用户状态
func (s *sUserManager) SetUserStatus(ctx context.Context, id uint64, status uint) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{Status: status}).Where(do.User{Id: id}).Update()
	return
}

// DeleteUser 删除用户
func (s *sUserManager) DeleteUser(ctx context.Context, ids []uint64) (err error) {
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		// 删除对应权限
		var enforcer *casbin.SyncedEnforcer
		enforcer, terr = Casbin(ctx).GetEnforcer()
		if terr != nil {
			return terr
		}
		for _, v := range ids {
			_, terr = enforcer.RemoveFilteredGroupingPolicy(0, gconv.String(v))
			if terr != nil {
				return terr
			}
		}
		return nil
	})
	return
}

// IsPassportAvailable 用户账号是否可用
func (s *sUserManager) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Unscoped().Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsMobileAvailable 用户手机是否可用
func (s *sUserManager) IsMobileAvailable(ctx context.Context, mobile string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Mobile: mobile}).Unscoped().Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfileList 获取个人中心信息列表
func (s *sUserManager) GetProfileList(ctx context.Context, userList []*entity.User) (profileInfos []*v1.UserProfileInfo, err error) {
	// 获取所有部门
	var allDepts []*entity.Dept
	if allDepts, err = Dept().GetAllDepts(ctx); err != nil {
		return
	}
	// 组装个人中心信息
	profileInfos = make([]*v1.UserProfileInfo, len(userList))
	for k, u := range userList {
		// 处理用户信息
		u.Password = ""
		u.Salt = ""
		profileInfos[k] = &v1.UserProfileInfo{
			User: u,
		}
		// 处理部门信息
		dept := Dept().GetDeptById(allDepts, u.DeptId)
		dept.DeletedAt = &gtime.Time{}
		var deptInfo *entity.Dept
		if err = gconv.Struct(gutil.Copy(dept), &deptInfo); err != nil {
			return
		}
		deptInfo.Name = Dept().GetDeptAllNameById(allDepts, u.DeptId)
		profileInfos[k].Dept = deptInfo
		// 处理角色信息
		var roles []*entity.Role
		if roles, err = Role().GetUserRoles(ctx, u.Id); err != nil {
			return
		}
		profileInfos[k].Roles = roles
	}
	return
}

// SearchByRealname 通过姓名搜索用户
func (s *sUserManager) SearchByRealname(ctx context.Context, realname string) (list []*entity.User, err error) {
	err = dao.User.Ctx(ctx).WhereLike(dao.User.Columns().Realname, "%"+realname+"%").Scan(&list)
	if err != nil {
		return
	}
	return
}

// saveUser 保存用户数据
func (s *sUserManager) saveUser(ctx context.Context, req *v1.UserAddReq, password, salt string) (userId int64, err error) {
	userId, err = dao.User.Ctx(ctx).Data(do.User{
		Passport: req.Passport,
		Password: password,
		Salt:     salt,
		Realname: req.Realname,
		Nickname: req.Nickname,
		DeptId:   req.DeptId,
		Gender:   req.Gender,
		Status:   req.Status,
		IsAdmin:  req.IsAdmin,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Remark:   req.Remark,
	}).FieldsEx(dao.User.Columns().Id).InsertAndGetId()
	return
}

// updateUser 更新用户数据
func (s *sUserManager) updateUser(ctx context.Context, req *v1.UserEditReq) (err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Realname: req.Realname,
		Nickname: req.Nickname,
		DeptId:   req.DeptId,
		Gender:   req.Gender,
		Status:   req.Status,
		IsAdmin:  req.IsAdmin,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Remark:   req.Remark,
	}).Where(do.User{Id: req.Id}).Update()
	return
}
