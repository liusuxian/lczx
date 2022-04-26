package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/internal/upload"
)

type sWdkService struct{}

var (
	insWdkService = sWdkService{}
)

// WdkService 文档库服务记录管理服务
func WdkService() *sWdkService {
	return &insWdkService
}

// GetWdkServiceRecord 获取文档库服务记录
func (s *sWdkService) GetWdkServiceRecord(ctx context.Context, projectId uint64) (list []*v1.WdkServiceRecordInfo, err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Where(do.WdkServiceRecord{ProjectId: projectId}).OrderDesc(dao.WdkServiceRecord.Columns().ServiceTime).
		ScanList(&list, "Record")
	if err != nil {
		return
	}
	err = dao.WdkServicePhoto.Ctx(ctx).Where(dao.WdkServicePhoto.Columns().Id, gdb.ListItemValuesUnique(list, "Record", "Id")).
		ScanList(&list, "Photo", "Record", "Id:Id")
	return
}

// AddWdkServiceRecord 新增文档库服务记录
func (s *sWdkService) AddWdkServiceRecord(ctx context.Context, req *v1.WdkServiceRecordAddReq, xch *upload.FileInfo, Photos []*upload.FileInfo) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 通过文档库项目ID判断文档库项目信息是否存在
		var terr error
		var wdkProjectExists bool
		wdkProjectExists, terr = WdkProject().WdkProjectExistsById(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		if !wdkProjectExists {
			terr = gerror.Newf(`文档库项目ID[%d]不存在`, req.ProjectId)
			return terr
		}
		// 写入文档库服务记录数据
		var recordId int64
		recordId, terr = dao.WdkServiceRecord.Ctx(ctx).Data(do.WdkServiceRecord{
			ProjectId:    req.ProjectId,
			ServiceTime:  req.ServiceTime,
			XchName:      xch.FileName,
			XchOriginUrl: xch.OriginFileUrl,
			XchPdfUrl:    xch.PdfFileUrl,
			Remark:       req.Remark,
		}).InsertAndGetId()
		if terr != nil {
			return terr
		}
		// 写入文档库服务记录照片数据
		photoData := g.List{}
		for _, file := range Photos {
			photoData = append(photoData, g.Map{
				"id":   recordId,
				"name": file.FileName,
				"url":  file.OriginFileUrl,
			})
		}
		_, terr = dao.WdkServicePhoto.Ctx(ctx).Data(photoData).Batch(len(g.List{})).Insert()
		return terr
	})
	return
}

// DeleteWdkServiceRecord 删除文档库服务记录
func (s *sWdkService) DeleteWdkServiceRecord(ctx context.Context, ids []uint64) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var terr error
		_, terr = dao.WdkServiceRecord.Ctx(ctx).WhereIn(dao.WdkServiceRecord.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		_, terr = dao.WdkServicePhoto.Ctx(ctx).WhereIn(dao.WdkServicePhoto.Columns().Id, ids).Delete()
		return terr
	})
	return
}
