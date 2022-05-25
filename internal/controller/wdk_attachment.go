package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
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
	// 上传附件
	var uploadAttachments []*ghttp.UploadFile
	if len(req.UploadAttachments) != 0 {
		uploadAttachments = req.UploadAttachments
	} else {
		uploadAttachments = []*ghttp.UploadFile{req.UploadAttachment}
	}
	var fileInfos []*upload.FileInfo
	fileInfos, err = upload.Upload.UploadFiles(uploadAttachments, "wdk/attachment")
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
