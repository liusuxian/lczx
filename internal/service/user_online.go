package service

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/grpool"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/utility/logger"
)

type sUserOnline struct {
	pool *grpool.Pool
}

var insUserOnline = sUserOnline{
	pool: grpool.New(100),
}

// UserOnline 用户在线表管理服务
func UserOnline() *sUserOnline {
	return &insUserOnline
}

// Invoke 异步保存用户在线状态
func (s *sUserOnline) Invoke(ctx context.Context, data *entity.UserOnline) {
	err := s.pool.Add(ctx, func(ctx context.Context) {
		// 保存用户在线状态
		s.SaveUserOnline(ctx, data)
	})
	if err != nil {
		logger.Error(ctx, "UserOnline Pool Add Error: ", err.Error())
	}
}

// SaveUserOnline 保存用户在线状态
func (s *sUserOnline) SaveUserOnline(ctx context.Context, data *entity.UserOnline) {
	// 查询是否已存在当前用户
	var info *entity.UserOnline
	var err error
	model := dao.UserOnline.Ctx(ctx)
	err = model.Fields(dao.UserOnline.Columns().Id).Where(do.UserOnline{Token: data.Token}).Scan(&info)
	if err != nil {
		logger.Error(ctx, "SaveUserOnline Error: ", err.Error())
	}
	if info != nil {
		// 已存在则更新
		_, err = model.Where(do.UserOnline{Id: info.Id}).FieldsEx(dao.UserOnline.Columns().Id).Update(data)
		if err != nil {
			logger.Error(ctx, "SaveUserOnline Update Error: ", err.Error())
		}
	} else {
		// 不存在则新增
		_, err = model.FieldsEx(dao.UserOnline.Columns().Id).Insert(data)
		if err != nil {
			logger.Error(ctx, "SaveUserOnline Insert Error: ", err.Error())
		}
	}
}

// DeleteOnlineByToken 删除用户在线状态操作
func (s *sUserOnline) DeleteOnlineByToken(ctx context.Context, token string) {
	_, err := dao.UserOnline.Ctx(ctx).Delete(do.UserOnline{Token: token})
	if err != nil {
		logger.Error(ctx, "DeleteOnlineByToken Error: ", err.Error())
	}
}

// GetOnlineList 获取在线用户列表
func (s *sUserOnline) GetOnlineList(ctx context.Context, req *v1.UserOnlineListReq) (total int, list []*entity.UserOnline, err error) {
	model := dao.UserOnline.Ctx(ctx)
	if req.Ip != "" {
		model = model.WhereLike(dao.UserOnline.Columns().Ip, "%"+req.Ip+"%")
	}
	if req.Passport != "" {
		model = model.WhereLike(dao.UserOnline.Columns().Passport, "%"+req.Passport+"%")
	}
	total, err = model.Count()
	if err != nil {
		err = gerror.Wrap(err, "获取总行数失败")
		return
	}
	err = model.FieldsEx(dao.UserOnline.Columns().Token).Page(req.CurPage, req.PageSize).
		OrderDesc(dao.UserOnline.Columns().Time).Scan(&list)
	if err != nil {
		err = gerror.Wrap(err, "获取数据失败")
		return
	}
	return
}

// ForceLogout 强退用户
func (s *sUserOnline) ForceLogout(ctx context.Context, ids []int) (tokens []string, err error) {
	var list []*entity.UserOnline
	list, err = s.GetOnlineListByIds(ctx, ids)
	if err != nil {
		return
	}
	for _, v := range list {
		tokens = append(tokens, v.Token)
	}
	_, err = dao.UserOnline.Ctx(ctx).WhereIn(dao.UserOnline.Columns().Id, ids).Delete()
	return
}

// GetOnlineListByIds 通过ID列表获取在线用户列表
func (s *sUserOnline) GetOnlineListByIds(ctx context.Context, ids []int) (list []*entity.UserOnline, err error) {
	err = dao.UserOnline.Ctx(ctx).WhereIn(dao.UserOnline.Columns().Id, ids).Scan(&list)
	return
}
