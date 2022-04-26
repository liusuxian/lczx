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
	WdkService = cWdkService{}
)

type cWdkService struct{}

// Record 文档库服务记录
func (c *cWdkService) Record(ctx context.Context, req *v1.WdkServiceRecordReq) (res *v1.WdkServiceRecordRes, err error) {
	var list []*v1.WdkServiceRecordInfo
	list, err = service.WdkService().GetWdkServiceRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkServiceRecordFailed, err)
		return
	}

	res = &v1.WdkServiceRecordRes{List: list}
	return
}

// Add 新增文档库服务记录
func (c *cWdkService) Add(ctx context.Context, req *v1.WdkServiceRecordAddReq) (res *v1.WdkServiceRecordAddRes, err error) {
	// 获取上传行程函文件信息
	xchFile := g.RequestFromCtx(ctx).GetUploadFile(req.XchUploadName)
	if xchFile == nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, gerror.New("获取上传行程函文件信息失败"))
		return
	}
	// 获取上传图片信息
	photoFiles := g.RequestFromCtx(ctx).GetUploadFiles(req.PhotoUploadName)
	if len(photoFiles) == 0 {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, gerror.New("获取上传图片信息失败"))
		return
	}
	if len(photoFiles) > 8 {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, gerror.New("最多可添加8张照片"))
		return
	}
	// 上传行程函文件
	var xch *upload.FileInfo
	xch, err = upload.Upload.UploadFile(xchFile, "wdk/service/xch")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, gerror.Wrap(err, "上传行程函文件失败"))
		return
	}
	// 上传图片
	var photos []*upload.FileInfo
	photos, err = upload.Upload.UploadFiles(photoFiles, "wdk/service/photo")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, gerror.Wrap(err, "上传图片失败"))
		return
	}
	// 新增文档库服务记录
	err = service.WdkService().AddWdkServiceRecord(ctx, req, xch, photos)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, err)
		return
	}

	return
}
