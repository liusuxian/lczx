package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
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
