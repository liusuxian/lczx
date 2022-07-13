package wdk

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/service"
)

type sWdkService struct{}

func init() {
	service.RegisterWdkService(newWdkService())
}

// 文档库服务记录管理服务
func newWdkService() *sWdkService {
	return &sWdkService{}
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
func (s *sWdkService) AddWdkService(ctx context.Context, req *v1.WdkServiceAddReq, xchFile *model.UploadFileInfo, photos []*model.UploadFileInfo) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 保存文档库服务记录
		var terr error
		terr = s.saveWdkServiceRecord(ctx, req, xchFile, photos)
		if terr != nil {
			return terr
		}
		// 设置文档库项目文件上传状态为正常
		terr = service.WdkProject().SetWdkProjectFileUploadStatusNormal(ctx, req.ProjectId)
		return terr
	})
	return
}

// DeleteWdkService 删除文档库服务记录
func (s *sWdkService) DeleteWdkService(ctx context.Context, ids []uint64) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除文档库服务记录
		var terr error
		_, terr = dao.WdkServiceRecord.Ctx(ctx).WhereIn(dao.WdkServiceRecord.Columns().Id, ids).Delete()
		if terr != nil {
			return terr
		}
		// 删除文档库服务照片
		_, terr = dao.WdkServicePhoto.Ctx(ctx).WhereIn(dao.WdkServicePhoto.Columns().Id, ids).Delete()
		return terr
	})
	return
}

// saveWdkServiceRecord 保存文档库服务记录
func (s *sWdkService) saveWdkServiceRecord(ctx context.Context, req *v1.WdkServiceAddReq, xchFile *model.UploadFileInfo, photos []*model.UploadFileInfo) (err error) {
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
