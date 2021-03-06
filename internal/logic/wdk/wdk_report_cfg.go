package wdk

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/internal/service"
)

type sWdkReportCfg struct{}

func init() {
	service.RegisterWdkReportCfg(newWdkReportCfg())
}

// 文档库报告类型配置管理服务
func newWdkReportCfg() *sWdkReportCfg {
	return &sWdkReportCfg{}
}

// GetWdkReportCfgList 获取文档库报告类型配置列表
func (s *sWdkReportCfg) GetWdkReportCfgList(ctx context.Context, name string) (list []*v1.WdkReportCfgInfo, err error) {
	model := dao.WdkReportCfg.Ctx(ctx)
	columns := dao.WdkReportCfg.Columns()
	if name != "" {
		model = model.WhereLike(columns.Name, "%"+name+"%")
	}
	err = model.OrderAsc(columns.Id).ScanList(&list, "ReportCfg")
	if err != nil {
		return
	}
	auditCfgModel := dao.WdkReportAuditCfg.Ctx(ctx)
	auditCfgColumns := dao.WdkReportAuditCfg.Columns()
	err = auditCfgModel.Where(auditCfgColumns.Id, gdb.ListItemValuesUnique(list, "ReportCfg", "Id")).
		ScanList(&list, "ReportAuditCfg", "ReportCfg", "Id:Id")
	return
}

// GetWdkAllReportCfg 获取文档库全部报告类型配置信息列表
func (s *sWdkReportCfg) GetWdkAllReportCfg(ctx context.Context) (list []*v1.WdkReportCfgInfo, err error) {
	err = dao.WdkReportCfg.Ctx(ctx).ScanList(&list, "ReportCfg")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditCfg.Ctx(ctx).Where(dao.WdkReportAuditCfg.Columns().Id, gdb.ListItemValuesUnique(list, "ReportCfg", "Id")).
		ScanList(&list, "ReportAuditCfg", "ReportCfg", "Id:Id")
	return
}

// AddWdkReportCfg 新增文档库报告类型配置
func (s *sWdkReportCfg) AddWdkReportCfg(ctx context.Context, req *v1.WdkReportCfgAddReq) (err error) {
	err = dao.WdkReportCfg.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查文档库报告类型名称是否可用
		var available bool
		var terr error
		available, terr = s.IsWdkReportCfgNameAvailable(ctx, req.Name)
		if terr != nil {
			return terr
		}
		if !available {
			return gerror.Newf(`文档库报告类型名称[%s]已存在`, req.Name)
		}
		// 写入文档库报告类型配置数据
		var typeId int64
		typeId, terr = dao.WdkReportCfg.Ctx(ctx).Data(do.WdkReportCfg{Name: req.Name}).
			FieldsEx(dao.WdkReportCfg.Columns().Id).InsertAndGetId()
		if terr != nil {
			return terr
		}
		// 处理客户端发送过来的文档库报告审核员用户ID列表
		auditUidSet := gset.NewFrom(req.AuditUids)
		var userList []*entity.User
		terr = dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, auditUidSet.Slice()).Scan(&userList)
		if terr != nil {
			return terr
		}
		if len(userList) == 0 {
			return gerror.Newf(`文档库报告审核员用户ID列表%v不存在`, req.AuditUids)
		}
		terr = s.saveWdkReportAuditCfg(ctx, userList, gconv.Uint64(typeId), req.Name)
		return terr
	})
	return
}

// GetWdkReportCfgById 通过文档库报告类型ID获取文档库报告类型配置信息
func (s *sWdkReportCfg) GetWdkReportCfgById(ctx context.Context, id uint64) (wdkReportCfg *v1.WdkReportCfgInfo, err error) {
	wdkReportCfg = &v1.WdkReportCfgInfo{}
	err = dao.WdkReportCfg.Ctx(ctx).Where(do.WdkReportCfg{Id: id}).Scan(&wdkReportCfg.ReportCfg)
	if err != nil {
		return
	}
	err = dao.WdkReportAuditCfg.Ctx(ctx).Where(do.WdkReportAuditCfg{Id: id}).Scan(&wdkReportCfg.ReportAuditCfg)
	return
}

// EditWdkReportCfg 编辑文档库报告类型配置
func (s *sWdkReportCfg) EditWdkReportCfg(ctx context.Context, req *v1.WdkReportCfgEditReq) (err error) {
	err = dao.WdkReportCfg.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查文档库报告类型配置是否存在
		var terr error
		var reportCfg *entity.WdkReportCfg
		terr = dao.WdkReportCfg.Ctx(ctx).Where(do.WdkReportCfg{Id: req.Id}).Scan(&reportCfg)
		if terr != nil {
			return terr
		}
		if reportCfg == nil {
			return gerror.Newf(`文档库报告类型ID[%d]不存在`, req.Id)
		}
		// 检查文档库报告类型名称是否可用
		if req.Name != reportCfg.Name {
			var available bool
			available, terr = s.IsWdkReportCfgNameAvailable(ctx, req.Name)
			if terr != nil {
				return terr
			}
			if !available {
				return gerror.Newf(`文档库报告类型名称[%s]已存在`, req.Name)
			}
		}
		// 更新文档库报告类型配置数据
		_, terr = dao.WdkReportCfg.Ctx(ctx).Data(do.WdkReportCfg{Name: req.Name}).Where(do.WdkReportCfg{Id: req.Id}).Update()
		if terr != nil {
			return terr
		}
		// 删除原来的审核配置数据
		_, terr = dao.WdkReportAuditCfg.Ctx(ctx).Where(do.WdkReportAuditCfg{Id: req.Id}).Delete()
		if terr != nil {
			return terr
		}
		// 处理客户端发送过来的文档库报告审核员用户ID列表
		auditUidSet := gset.NewFrom(req.AuditUids)
		var userList []*entity.User
		terr = dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, auditUidSet.Slice()).Scan(&userList)
		if terr != nil {
			return terr
		}
		if len(userList) == 0 {
			return gerror.Newf(`文档库报告审核员用户ID列表%v不存在`, req.AuditUids)
		}
		terr = s.saveWdkReportAuditCfg(ctx, userList, req.Id, req.Name)
		return terr
	})
	return
}

// DeleteWdkReportCfg 删除文档库报告类型配置
func (s *sWdkReportCfg) DeleteWdkReportCfg(ctx context.Context, ids []uint64) (err error) {
	err = dao.WdkReportCfg.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.WdkReportCfg.Ctx(ctx).WhereIn(dao.WdkReportCfg.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		_, terr = dao.WdkReportAuditCfg.Ctx(ctx).WhereIn(dao.WdkReportAuditCfg.Columns().Id, ids).Delete()
		return terr
	})
	return
}

// IsWdkReportCfgNameAvailable 文档库报告类型名称是否可用
func (s *sWdkReportCfg) IsWdkReportCfgNameAvailable(ctx context.Context, name string) (bool, error) {
	count, err := dao.WdkReportCfg.Ctx(ctx).Where(do.WdkReportCfg{Name: name}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetWdkReportCfgCount 获取文档库报告类型配置数量
func (s *sWdkReportCfg) GetWdkReportCfgCount(ctx context.Context) (count int, err error) {
	count, err = dao.WdkReportCfg.Ctx(ctx).Count()
	return
}

// GetWdkReportCfgByIds 通过报告类型ID列表获取报告类型配置
func (s *sWdkReportCfg) GetWdkReportCfgByIds(ctx context.Context, ids []uint64) (reportCfgInfos []*v1.WdkReportCfgInfo, err error) {
	err = dao.WdkReportCfg.Ctx(ctx).WhereIn(dao.WdkReportCfg.Columns().Id, ids).OrderAsc(dao.WdkReportCfg.Columns().Id).
		ScanList(&reportCfgInfos, "ReportCfg")
	if err != nil {
		return
	}
	err = dao.WdkReportAuditCfg.Ctx(ctx).Where(dao.WdkReportAuditCfg.Columns().Id, gdb.ListItemValuesUnique(reportCfgInfos, "ReportCfg", "Id")).
		ScanList(&reportCfgInfos, "ReportAuditCfg", "ReportCfg", "Id:Id")
	return
}

// saveWdkReportAuditCfg 保存文档库报告审核配置
func (s *sWdkReportCfg) saveWdkReportAuditCfg(ctx context.Context, userList []*entity.User, id uint64, name string) (err error) {
	auditUserData := g.List{}
	for _, user := range userList {
		auditUserData = append(auditUserData, g.Map{
			"id":         id,
			"audit_uid":  user.Id,
			"type_name":  name,
			"audit_name": user.Realname,
		})
	}
	_, err = dao.WdkReportAuditCfg.Ctx(ctx).Data(auditUserData).Batch(len(auditUserData)).Insert()
	return
}
