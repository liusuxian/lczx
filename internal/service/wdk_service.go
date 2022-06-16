package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
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
func (s *sWdkService) GetWdkServiceRecord(ctx context.Context, projectId uint64) (list []*v1.WdkServiceInfo, err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Where(do.WdkServiceRecord{ProjectId: projectId}).
		OrderDesc(dao.WdkServiceRecord.Columns().ServiceTime).ScanList(&list, "Record")
	if err != nil {
		return
	}
	err = dao.WdkServicePhoto.Ctx(ctx).Where(dao.WdkServicePhoto.Columns().Id, gdb.ListItemValuesUnique(list, "Record", "Id")).
		ScanList(&list, "Photo", "Record", "Id:Id")
	return
}

// AddWdkService 新增文档库服务记录
func (s *sWdkService) AddWdkService(ctx context.Context, req *v1.WdkServiceAddReq, xchFile *upload.FileInfo, photos []*upload.FileInfo) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 保存文档库服务记录
		return s.saveWdkServiceRecord(ctx, req, xchFile, photos)
	})
	return
}

// saveWdkServiceRecord 保存文档库服务记录
func (s *sWdkService) saveWdkServiceRecord(ctx context.Context, req *v1.WdkServiceAddReq, xchFile *upload.FileInfo, photos []*upload.FileInfo) (err error) {
	var recordId int64
	recordId, err = dao.WdkServiceRecord.Ctx(ctx).Data(do.WdkServiceRecord{
		ProjectId:    req.ProjectId,
		ServiceTime:  req.ServiceTime,
		XchName:      xchFile.FileName,
		XchOriginUrl: xchFile.OriginFileUrl,
		XchPdfUrl:    xchFile.PdfFileUrl,
		Remark:       req.Remark,
	}).FieldsEx(dao.WdkServiceRecord.Columns().Id).InsertAndGetId()
	if err != nil {
		return
	}
	// 保存文档库服务照片数据
	photoData := g.List{}
	for _, file := range photos {
		photoData = append(photoData, g.Map{
			"id":   recordId,
			"name": file.FileName,
			"url":  file.OriginFileUrl,
		})
	}
	_, err = dao.WdkServicePhoto.Ctx(ctx).Data(photoData).Batch(len(photoData)).Insert()
	return
}
