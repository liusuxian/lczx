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
func (s *sWdkService) AddWdkService(ctx context.Context, req *v1.WdkServiceAddReq, xch *upload.FileInfo, photos []*upload.FileInfo) (err error) {
	err = dao.WdkServiceRecord.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查新增文档库服务记录权限
		var terr error
		var wdkProject *v1.WdkProjectInfo
		wdkProject, terr = s.AuthAdd(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		// 保存文档库服务记录
		terr = s.saveWdkServiceRecord(ctx, req, xch, photos)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目阶段
		terr = WdkProject().SetWdkProjectStep(ctx, wdkProject.ProjectInfo.Id, 7)
		return terr
	})
	return
}

// AuthAdd 检查新增文档库服务记录权限
func (s *sWdkService) AuthAdd(ctx context.Context, projectId uint64) (wdkProject *v1.WdkProjectInfo, err error) {
	// 通过文档库项目ID判断文档库项目信息是否存在
	wdkProject, err = WdkProject().GetWdkProjectById(ctx, projectId)
	if err != nil {
		return
	}
	if wdkProject == nil || wdkProject.ProjectInfo == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, projectId)
		return
	}
	// 判断写入权限
	user := Context().Get(ctx).User
	if user.Id != wdkProject.ProjectInfo.PrincipalUid {
		err = gerror.New("抱歉！！！该项目您没有添加服务记录的权限")
		return
	}
	return
}

// saveWdkServiceRecord 保存文档库服务记录
func (s *sWdkService) saveWdkServiceRecord(ctx context.Context, req *v1.WdkServiceAddReq, xch *upload.FileInfo, photos []*upload.FileInfo) (err error) {
	var recordId int64
	recordId, err = dao.WdkServiceRecord.Ctx(ctx).Data(do.WdkServiceRecord{
		ProjectId:    req.ProjectId,
		ServiceTime:  req.ServiceTime,
		XchName:      xch.FileName,
		XchOriginUrl: xch.OriginFileUrl,
		XchPdfUrl:    xch.PdfFileUrl,
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
