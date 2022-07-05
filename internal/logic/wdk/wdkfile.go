package wdk

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/dao"
	"lczx/internal/model"
	"lczx/internal/model/do"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/internal/upload"
)

type sWdkFile struct{}

func init() {
	service.RegisterWdkFile(newWdkFile())
}

// 文档库上传文件记录管理服务
func newWdkFile() *sWdkFile {
	return &sWdkFile{}
}

// GetWdkFileRecord 获取文档库上传文件记录
func (s *sWdkFile) GetWdkFileRecord(ctx context.Context, projectId uint64) (list []*entity.WdkFile, err error) {
	err = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: projectId}).OrderAsc(dao.WdkFile.Columns().Type).Scan(&list)
	return
}

// AddWdkFile 新增文档库上传文件记录
func (s *sWdkFile) AddWdkFile(ctx context.Context, req *v1.WdkFileAddReq, fileInfos []*upload.FileInfo) (err error) {
	err = dao.WdkFile.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 删除旧文件
		var terr error
		_, terr = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: req.ProjectId, Type: req.Type}).Delete()
		if terr != nil {
			return terr
		}
		// 保存文档库上传文件数据
		terr = s.saveWdkFile(ctx, req, fileInfos)
		if terr != nil {
			return terr
		}
		// 设置所属文档库项目阶段
		terr = service.WdkProject().SetWdkProjectStepByFileType(ctx, req.ProjectId, req.Type)
		if terr != nil {
			return terr
		}
		// 设置文档库项目文件上传状态为已完成
		terr = service.WdkProject().SetWdkProjectFileUploadStatusFinish(ctx, req.ProjectId)
		return terr
	})
	return
}

// GetWdkFileCountByProjectId 通过项目ID获取文档库项目上传文件类型数量
func (s *sWdkFile) GetWdkFileCountByProjectId(ctx context.Context, projectId uint64) (count int, err error) {
	count, err = dao.WdkFile.Ctx(ctx).Fields(dao.WdkFile.Columns().Type).Where(do.WdkFile{ProjectId: projectId}).
		Distinct().Count()
	return
}

// saveWdkFile 保存文档库上传文件数据
func (s *sWdkFile) saveWdkFile(ctx context.Context, req *v1.WdkFileAddReq, fileInfos []*upload.FileInfo) (err error) {
	var user *model.ContextUser
	user, err = service.Context().GetUser(ctx)
	if err != nil {
		return
	}
	wdkFileData := g.List{}
	for _, file := range fileInfos {
		wdkFileData = append(wdkFileData, g.Map{
			"project_id":  req.ProjectId,
			"name":        file.FileName,
			"type":        req.Type,
			"create_by":   user.Id,
			"create_name": user.Realname,
			"origin_url":  file.OriginFileUrl,
			"pdf_url":     file.PdfFileUrl,
		})
	}
	_, err = dao.WdkFile.Ctx(ctx).Data(wdkFileData).Batch(len(wdkFileData)).Insert()
	return
}
