package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/model/entity"
	"lczx/internal/service/internal/dao"
	"lczx/internal/service/internal/do"
	"lczx/internal/upload"
)

type sWdkFile struct{}

var (
	insWdkFile = sWdkFile{}
)

// WdkFile 文档库上传文件记录管理服务
func WdkFile() *sWdkFile {
	return &insWdkFile
}

// GetWdkFileRecord 获取文档库上传文件记录
func (s *sWdkFile) GetWdkFileRecord(ctx context.Context, projectId uint64) (list []*entity.WdkFile, err error) {
	err = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: projectId}).OrderAsc(dao.WdkFile.Columns().Type).Scan(&list)
	return
}

// AddWdkFile 新增文档库上传文件记录
func (s *sWdkFile) AddWdkFile(ctx context.Context, req *v1.WdkFileAddReq, file *upload.FileInfo) (err error) {
	err = dao.WdkFile.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 检查新增文档库上传文件记录权限
		var terr error
		terr = s.AuthAdd(ctx, req.ProjectId)
		if terr != nil {
			return terr
		}
		// 获取上传文件类型是否已存在
		var wdkFileInfo *entity.WdkFile
		terr = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: req.ProjectId, Type: req.Type}).Scan(&wdkFileInfo)
		if terr != nil {
			return terr
		}
		// 不存在则新增
		if wdkFileInfo == nil {
			// 写入文档库上传文件记录数据
			user := Context().Get(ctx).User
			_, terr = dao.WdkFile.Ctx(ctx).Data(do.WdkFile{
				ProjectId:  req.ProjectId,
				Name:       file.FileName,
				Type:       req.Type,
				CreateBy:   user.Id,
				CreateName: user.Realname,
				OriginUrl:  file.OriginFileUrl,
				PdfUrl:     file.PdfFileUrl,
			}).Insert()
		} else {
			// 存在则更新
			_, terr = dao.WdkFile.Ctx(ctx).Data(do.WdkFile{
				Name:      file.FileName,
				OriginUrl: file.OriginFileUrl,
				PdfUrl:    file.PdfFileUrl,
			}).Where(do.WdkFile{
				ProjectId: wdkFileInfo.ProjectId,
				Type:      wdkFileInfo.Type,
			}).Update()
		}
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目阶段
		terr = WdkProject().SetWdkProjectStep(ctx, req.ProjectId, req.Type)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目文件上传状态为是
		terr = WdkProject().SetWdkProjectFileUploadStatus(ctx, req.ProjectId)
		return terr
	})
	return
}

// AuthAdd 检查新增文档库上传文件记录权限
func (s *sWdkFile) AuthAdd(ctx context.Context, projectId uint64) (err error) {
	// 通过文档库项目ID判断文档库项目信息是否存在
	var wdkProject *entity.WdkProject
	wdkProject, err = WdkProject().GetWdkProjectById(ctx, projectId)
	if err != nil {
		return
	}
	if wdkProject == nil {
		err = gerror.Newf(`文档库项目ID[%d]不存在`, projectId)
		return
	}
	// 判断写入权限
	user := Context().Get(ctx).User
	if user.Id != wdkProject.PrincipalUid && user.IsAdmin != 1 {
		err = gerror.New("抱歉！！！该项目您没有上传文件的权限")
		return
	}
	return
}

// GetWdkFileCountByProjectId 通过项目ID获取文档库项目上传文件记录数量
func (s *sWdkFile) GetWdkFileCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	count, err = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: projectId}).Count()
	return
}
