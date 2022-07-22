package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/model"
	"lczx/internal/service"
)

var (
	WdkReport = cWdkReport{}
)

type cWdkReport struct{}

// Record 文档库上传报告记录
func (c *cWdkReport) Record(ctx context.Context, req *v1.WdkReportRecordReq) (res *v1.WdkReportRecordRes, err error) {
	var list []*v1.WdkReportInfo
	list, err = service.WdkReport().GetWdkReportRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportRecordFailed, err)
		return
	}

	res = &v1.WdkReportRecordRes{List: list}
	return
}

// Add 新增文档库上传报告记录
func (c *cWdkReport) Add(ctx context.Context, req *v1.WdkReportAddReq) (res *v1.WdkReportAddRes, err error) {
	// 上传报告
	var report *model.UploadFileInfo
	report, err = service.AliyunOSS().UploadFile(req.UploadReport, "wdk/report")
	if err != nil {
		err = gerror.WrapCode(code.AddWdkReportRecordFailed, err)
		return
	}
	// 新增文档库上传报告记录
	err = service.WdkReport().AddWdkReport(ctx, req, report)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkReportRecordFailed, err)
		return
	}

	return
}

// Delete 删除文档库上传报告记录
func (c *cWdkReport) Delete(ctx context.Context, req *v1.WdkReportDeleteReq) (res *v1.WdkReportDeleteRes, err error) {
	err = service.WdkReport().DeleteWdkReport(ctx, req.Ids)
	if err != nil {
		err = gerror.WrapCode(code.DeleteWdkReportRecordFailed, err)
		return
	}

	return
}

// ClientOptions 客户端选项
func (c *cWdkReport) ClientOptions(ctx context.Context, req *v1.WdkReportClientOptionsReq) (res *v1.WdkReportClientOptionsRes, err error) {
	// 获取客户端选项Map
	clientOptionMap := service.WdkReport().GetClientOptionMap()

	res = &v1.WdkReportClientOptionsRes{
		ExcellenceStatusList: clientOptionMap["excellenceStatusList"],
	}
	return
}

// List 文档库报告列表
func (c *cWdkReport) List(ctx context.Context, req *v1.WdkReportListReq) (res *v1.WdkReportListRes, err error) {
	var total int
	var list []*v1.WdkReportInfo
	total, list, err = service.WdkReport().GetWdkReportList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportListFailed, err)
		return
	}

	res = &v1.WdkReportListRes{
		CurPage: req.CurPage,
		Total:   total,
		List:    list,
	}
	return
}

// ChooseExcellence 文档库报告评选优秀报告
func (c *cWdkReport) ChooseExcellence(ctx context.Context, req *v1.WdkReportChooseExcellenceReq) (res *v1.WdkReportChooseExcellenceRes, err error) {
	err = service.WdkReport().SetWdkReportExcellence(ctx, req.Id, req.Excellence)
	if err != nil {
		err = gerror.WrapCode(code.WdkReportChooseExcellenceFailed, err)
		return
	}

	return
}

// ReportDownload 文档库报告文件下载
func (c *cWdkReport) ReportDownload(ctx context.Context, req *v1.WdkReportDownloadReq) (res *v1.WdkReportDownloadRes, err error) {
	var fileUrl string
	fileUrl, err = service.AliyunOSS().AuthorizationDownload(req.FilePath)
	if err != nil {
		err = gerror.WrapCode(code.WdkReportDownloadFileFailed, err)
		return
	}

	res = &v1.WdkReportDownloadRes{FileUrl: fileUrl}
	return
}
