package controller

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
	"lczx/internal/upload"
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
	// 检查新增文档库上传报告记录权限
	_, err = service.WdkReport().AuthAdd(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.AddWdkReportRecordFailed, err)
		return
	}
	// 获取上传报告信息
	file := g.RequestFromCtx(ctx).GetUploadFile(req.UploadName)
	if file == nil {
		err = gerror.WrapCode(code.AddWdkReportRecordFailed, gerror.New("获取上传报告信息失败"))
		return
	}
	// 上传报告
	var report *upload.FileInfo
	report, err = upload.Upload.UploadFile(file, "wdk/report")
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

// ExcellenceList 文档库优秀报告列表
func (c *cWdkReport) ExcellenceList(ctx context.Context, req *v1.WdkReportExcellenceListReq) (res *v1.WdkReportExcellenceListRes, err error) {
	var total int
	var list []*v1.WdkReportInfo
	total, list, err = service.WdkReport().GetWdkReportExcellenceList(ctx, req)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkReportExcellenceListFailed, err)
		return
	}

	res = &v1.WdkReportExcellenceListRes{
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
