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
	err = dao.WdkFile.Ctx(ctx).Where(do.WdkFile{ProjectId: projectId}).Scan(&list)
	return
}

//
func (s *sWdkFile) AddWdkFile(ctx context.Context, req *v1.WdkFileAddReq, file *upload.FileInfo) (err error) {
	err = dao.WdkFile.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		//
		return nil
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
