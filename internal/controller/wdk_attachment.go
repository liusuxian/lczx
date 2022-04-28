package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
	"lczx/internal/upload"
)

var (
	WdkAttachment = cWdkAttachment{}
)

type cWdkAttachment struct{}

// Record 文档库上传附件记录
func (c *cWdkAttachment) Record(ctx context.Context, req *v1.WdkAttachmentRecordReq) (res *v1.WdkAttachmentRecordRes, err error) {
	var list []*v1.WdkAttachmentInfo
	list, err = service.WdkAttachment().GetWdkAttachmentRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkAttachmentRecordFailed, err)
		return
	}

	res = &v1.WdkAttachmentRecordRes{List: list}
	return
}

// Add 新增文档库上传附件记录
func (c *cWdkAttachment) Add(ctx context.Context, req *v1.WdkAttachmentAddReq) (res *v1.WdkAttachmentAddRes, err error) {
	// 检查新增文档库上传附件记录权限
	err = service.WdkAttachment().AuthAdd(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkAttachmentRecordFailed, err)
		return
	}
	// 获取上传附件信息
	files := g.RequestFromCtx(ctx).GetUploadFiles(req.UploadName)
	if len(files) == 0 {
		err = gerror.WrapCode(code.AddWdkAttachmentRecordFailed, gerror.New("获取上传附件信息失败"))
		return
	}
	if len(files) > 4 {
		err = gerror.WrapCode(code.AddWdkAttachmentRecordFailed, gerror.New("最多可添加4份附件"))
		return
	}
	// 上传附件
	var fileInfos []*upload.FileInfo
	fileInfos, err = upload.Upload.UploadFiles(files, "wdk/attachment")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkAttachmentRecordFailed, err)
		return
	}
	// 新增文档库上传附件记录
	err = service.WdkAttachment().AddWdkAttachment(ctx, req, fileInfos)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkAttachmentRecordFailed, err)
		return
	}

	return
}
