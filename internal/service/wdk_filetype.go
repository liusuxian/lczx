package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
)

type sWdkFiletype struct{}

var (
	insWdkFiletype = sWdkFiletype{}
)

// WdkFiletype 文档库上传文件类型管理服务
func WdkFiletype() *sWdkFiletype {
	return &insWdkFiletype
}

// GetWdkFiletypeList 获取文档库上传文件类型列表
func (s *sWdkFiletype) GetWdkFiletypeList(ctx context.Context, req *v1.WdkFiletypeListReq) (total int, list []*v1.WdkFiletypeInfo, err error) {
	model := dao.WdkFiletypeCfg.Ctx(ctx)
	columns := dao.WdkFiletypeCfg.Columns()
	if req.Name != "" {
		model = model.WhereLike(columns.Name, "%"+req.Name+"%")
	}
	if req.Multiple != "" {
		model = model.Where(columns.Multiple, gconv.Uint(req.Multiple))
	}
	if req.Audit != "" {
		model = model.Where(columns.Audit, gconv.Uint(req.Audit))
	}
	if req.Step != "" {
		model = model.Where(columns.Step, gconv.Uint(req.Step))
	}
	if req.StartTime.String() != "" {
		model = model.WhereGTE(columns.CreateAt, req.StartTime)
	}
	if req.EndTime.String() != "" {
		model = model.WhereLTE(columns.CreateAt, req.EndTime)
	}
	total, err = model.Count()
	if err != nil {
		return
	}
	err = model.Page(req.CurPage, req.PageSize).OrderAsc(columns.Id).ScanList(&list, "FiletypeCfg")
	if err != nil {
		return
	}
	auditCfgModel := dao.WdkAuditCfg.Ctx(ctx)
	auditCfgColumns := dao.WdkAuditCfg.Columns()
	err = auditCfgModel.Where(auditCfgColumns.TypeId, gdb.ListItemValuesUnique(list, "FiletypeCfg", "Id")).
		ScanList(&list, "AuditCfg", "FiletypeCfg", "TypeId:Id")
	return
}

// AddWdkFiletype 新增文档库上传文件类型
func (s *sWdkFiletype) AddWdkFiletype(ctx context.Context, req *v1.WdkFiletypeAddReq) (err error) {
	err = dao.WdkFiletypeCfg.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查文档库上传文件类型名称是否可用
		var available bool
		var terr error
		available, terr = s.IsWdkFiletypeNameAvailable(ctx, req.Name)
		if terr != nil {
			return terr
		}
		if !available {
			terr = gerror.Newf(`文档库上传文件类型名称[%s]已存在`, req.Name)
			return terr
		}
		// 写入文档库上传文件类型数据
		var typeId int64
		typeId, terr = dao.WdkFiletypeCfg.Ctx(ctx).Data(do.WdkFiletypeCfg{
			Name:     req.Name,
			Multiple: req.Multiple,
			Audit:    req.Audit,
			Step:     req.Step,
		}).InsertAndGetId()
		if terr != nil {
			return terr
		}
		// 是否需要审核
		if req.Audit == 1 {
			// 处理客户端发送过来的文档库报告审核员用户ID列表
			auditUidsSet := gset.NewFrom(req.AuditUids)
			auditUserData := g.List{}
			for _, auid := range auditUidsSet.Slice() {
				var user *entity.User
				user, terr = User().GetUserById(ctx, gconv.Uint64(auid))
				if terr != nil {
					return terr
				}
				if user == nil {
					return gerror.Newf(`文档库报告审核员用户ID[%d]不存在`, auid)
				}
				auditUserData = append(auditUserData, g.Map{"type_id": typeId, "audit_uid": user.Id, "audit_name": user.Realname})
			}
			_, terr = dao.WdkAuditCfg.Ctx(ctx).Data(auditUserData).Batch(len(g.List{})).Insert()
			return terr
		}
		return nil
	})
	return
}

// IsWdkFiletypeNameAvailable 文档库上传文件类型名称是否可用
func (s *sWdkFiletype) IsWdkFiletypeNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.WdkFiletypeCfg.Ctx(ctx).Where(do.WdkFiletypeCfg{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
