package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model"
	"lczx/internal/service"
)

var (
	WdkService = cWdkService{}
)

type cWdkService struct{}

// Record 文档库服务记录
func (c *cWdkService) Record(ctx context.Context, req *v1.WdkServiceRecordReq) (res *v1.WdkServiceRecordRes, err error) {
	var list []*v1.WdkServiceInfo
	list, err = service.WdkService().GetWdkServiceRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkServiceRecordFailed, err)
		return
	}
	// 获取照片访问url
	for _, item := range list {
		for _, photo := range item.Photo {
			photo.Url, err = service.AliyunOSS().GetAccessUrl(ctx, photo.Url)
			if err != nil {
				err = gerror.WrapCode(code.GetWdkServiceRecordFailed, err)
				return
			}
		}
	}
	res = &v1.WdkServiceRecordRes{List: list}
	return
}

// Add 新增文档库服务记录
func (c *cWdkService) Add(ctx context.Context, req *v1.WdkServiceAddReq) (res *v1.WdkServiceAddRes, err error) {
	// 上传行程函文件
	var xchFile *model.UploadFileInfo
	xchFile, err = service.AliyunOSS().UploadFile(req.UploadXchFile, "wdk/service/xch")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, err)
		return
	}
	// 上传图片
	var uploadPhotos []*ghttp.UploadFile
	if len(req.UploadPhotos) != 0 {
		uploadPhotos = req.UploadPhotos
	} else {
		uploadPhotos = []*ghttp.UploadFile{req.UploadPhoto}
	}
	var photos []*model.UploadFileInfo
	photos, err = service.AliyunOSS().UploadImgs(uploadPhotos, "wdk/service/photo")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, err)
		return
	}
	// 新增文档库服务记录
	err = service.WdkService().AddWdkService(ctx, req, xchFile, photos)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkServiceRecordFailed, err)
		return
	}

	return
}

// Delete 删除文档库服务记录
func (c *cWdkService) Delete(ctx context.Context, req *v1.WdkServiceDeleteReq) (res *v1.WdkServiceDeleteRes, err error) {
	err = service.WdkService().DeleteWdkService(ctx, req.Ids, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.DeleteWdkServiceRecordFailed, err)
		return
	}

	return
}
