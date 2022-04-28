package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/internal/upload"
)

var (
	WdkFile = cWdkFile{}
)

type cWdkFile struct{}

// Record 文档库上传文件记录
func (c *cWdkFile) Record(ctx context.Context, req *v1.WdkFileRecordReq) (res *v1.WdkFileRecordRes, err error) {
	var list []*entity.WdkFile
	list, err = service.WdkFile().GetWdkFileRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkFileRecordFailed, err)
		return
	}

	res = &v1.WdkFileRecordRes{List: list}
	return
}

// Add 新增文档库上传文件记录
func (c *cWdkFile) Add(ctx context.Context, req *v1.WdkFileAddReq) (res *v1.WdkFileAddRes, err error) {
	// 检查新增文档库上传文件记录权限
	err = service.WdkFile().AuthAdd(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkFileRecordFailed, err)
		return
	}
	// 获取上传文件信息
	file := g.RequestFromCtx(ctx).GetUploadFile(req.UploadName)
	if file == nil {
		err = gerror.WrapCode(code.AddWdkFileRecordFailed, gerror.New("获取上传文件信息失败"))
		return
	}
	// 上传文件
	var fileInfo *upload.FileInfo
	fileInfo, err = upload.Upload.UploadFile(file, "wdk/file")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkFileRecordFailed, err)
		return
	}
	// 新增文档库上传文件记录
	err = service.WdkFile().AddWdkFile(ctx, req, fileInfo)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkFileRecordFailed, err)
		return
	}

	return
}
