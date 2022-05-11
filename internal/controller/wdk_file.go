package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model/entity"
	"lczx/internal/service"
	"lczx/internal/upload"
	"lczx/utility/response"
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
func (c *cWdkFile) Add(req *ghttp.Request) {
	ctx := req.GetCtx()
	var err error
	var wdkFileAddReq *v1.WdkFileAddReq
	if err = req.Parse(&wdkFileAddReq); err != nil {
		errCode := err.(gvalid.Error).Code().Code()
		errMsg := err.(gvalid.Error).Current().Error()
		response.RespJsonExit(req, errCode, errMsg)
	}
	// 检查新增文档库上传文件记录权限
	_, err = service.WdkFile().AuthAdd(ctx, wdkFileAddReq.ProjectId)
	if err != nil {
		response.RespJsonExit(req, code.AddWdkFileRecordFailed.Code(), code.AddWdkFileRecordFailed.Message()+": "+err.Error())
	}
	// 获取上传文件信息
	file := req.GetUploadFile("upload")
	if file == nil {
		response.RespJsonExit(req, code.AddWdkFileRecordFailed.Code(), code.AddWdkFileRecordFailed.Message()+": 获取上传文件信息失败")
	}
	// 上传文件
	var fileInfo *upload.FileInfo
	fileInfo, err = upload.Upload.UploadFile(file, "wdk/file")
	if err != nil {
		response.RespJsonExit(req, code.AddWdkFileRecordFailed.Code(), code.AddWdkFileRecordFailed.Message()+": "+err.Error())
	}
	// 新增文档库上传文件记录
	err = service.WdkFile().AddWdkFile(ctx, wdkFileAddReq, fileInfo)
	if err != nil {
		response.RespJsonExit(req, code.AddWdkFileRecordFailed.Code(), code.AddWdkFileRecordFailed.Message()+": "+err.Error())
	}

	response.Succ(req, v1.WdkFileAddRes{FileInfo: fileInfo})
	return
}
