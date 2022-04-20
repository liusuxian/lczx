package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
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
	deptIdsMap := gmap.New(true)
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
		model = model.Where(columns.Status, gconv.Uint(req.Status))
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.CreateAt, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.CreateAt, req.EndTime)
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
	// 检查用户账号是否可用
	var available bool
	available, err = s.IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`账号[%s]已存在`, req.Passport)
		return
	}
	// 检查用户手机是否可用
	available, err = s.IsMobileAvailable(ctx, req.Mobile)
	if err != nil {
		return
	}
	if !available {
		err = gerror.Newf(`手机[%s]已存在`, req.Mobile)
		return
	}
	// 通过部门ID判断部门信息是否存在
	var deptExists bool
	deptExists, err = Dept().DeptExistsById(ctx, req.DeptId)
	if err != nil {
		return
	}
	if !deptExists {
		err = gerror.Newf(`部门ID[%d]不存在`, req.DeptId)
		return
	}
	// 生成随机密码盐
	salt := grand.S(10)
	// 密码加密
	password := utils.EncryptPassword(req.Password, salt)
	// 开启事务
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 写入用户数据
		var userId int64
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
		if err != nil {
			return err
		}
		// 获取全部可用的角色
		var enableRoles []*entity.Role
		enableRoles, err = Role().GetEnableRoles(ctx)
		if err != nil {
			return err
		}
		roleIdsMap := gmap.New(true)
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
		err = Role().AddUserRoles(ctx, req.RoleIds, gconv.Uint64(userId))
		return err
	})
	return
}

// EditUser 编辑用户
func (s *sUserManager) EditUser(ctx context.Context, req *v1.UserEditReq) (err error) {
	// 检查用户手机是否可用
	var user *entity.User
	user, err = User().GetUserById(ctx, req.Id)
	if err != nil {
		return
	}
	if user == nil {
		err = gerror.Newf(`用户ID[%d]不存在`, req.Id)
		return
	}
	if user.Mobile != req.Mobile {
		var available bool
		available, err = s.IsMobileAvailable(ctx, req.Mobile)
		if err != nil {
			return
		}
		if !available {
			err = gerror.Newf(`手机[%s]已存在`, req.Mobile)
			return
		}
	}
	// 通过部门ID判断部门信息是否存在
	var deptExists bool
	deptExists, err = Dept().DeptExistsById(ctx, req.DeptId)
	if err != nil {
		return
	}
	if !deptExists {
		err = gerror.Newf(`部门ID[%d]不存在`, req.DeptId)
		return
	}
	// 开启事务
	err = dao.User.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 更新用户数据
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
		if err != nil {
			return err
		}
		// 获取全部可用的角色
		var enableRoles []*entity.Role
		enableRoles, err = Role().GetEnableRoles(ctx)
		if err != nil {
			return err
		}
		roleIdsMap := gmap.New(true)
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
		err = Role().EditUserRoles(ctx, req.RoleIds, req.Id)
		return err
	})
	if err != nil {
		return
	}
	// 清除用户缓存
	_, err = Cache().ClearCache(ctx, User().UserCacheKey(req.Id))
	return
}

// ResetUserPwd 重置用户密码
func (s *sUserManager) ResetUserPwd(ctx context.Context, id uint64, newPassword string) (err error) {
	// 生成随机密码盐
	salt := grand.S(10)
	newEncryptPassword := utils.EncryptPassword(newPassword, salt)
	_, err = dao.User.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     User().UserCacheKey(id),
		Force:    false,
	}).Data(do.User{
		Salt:     salt,
		Password: newEncryptPassword,
	}).Where(do.User{Id: id}).Update()
	return
}

// IsPassportAvailable 用户账号是否可用
func (s *sUserManager) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Passport: passport}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsMobileAvailable 用户手机是否可用
func (s *sUserManager) IsMobileAvailable(ctx context.Context, mobile string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{Mobile: mobile}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfileList 获取个人中心信息列表
func (s *sUserManager) GetProfileList(ctx context.Context, userList []*entity.User) (profileInfos []*v1.UserProfileInfo, err error) {
	// 获取所有部门
	var allDepts []*entity.Dept
	allDepts, err = Dept().GetAllDepts(ctx)
	if err != nil {
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
		deptNames := Dept().GetDeptAllNameById(allDepts, u.DeptId)
		utils.Reverse(deptNames)
		dept := Dept().GetDeptById(allDepts, u.DeptId)
		dept.Name = gstr.Join(deptNames, "/")
		profileInfos[k].Dept = dept
		// 处理角色信息
		var roles []*entity.Role
		roles, err = Role().GetUserRoles(ctx, u.Id)
		if err != nil {
			return
		}
		profileInfos[k].Roles = roles
	}
	return
}
