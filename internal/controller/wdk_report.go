package controller

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/net/context"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
	"lczx/internal/upload"
	"lczx/utility/response"
	"strings"
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
func (c *cWdkReport) Add(req *ghttp.Request) {
	ctx := req.GetCtx()
	var err error
	paramMap := req.GetMap()
	projectId := gconv.Uint64(paramMap["projectId"])
	tmpTypeIds := strings.Split(gconv.String(paramMap["typeIds"]), ",")
	typeIds := make([]uint64, 0, len(tmpTypeIds))
	for _, v := range tmpTypeIds {
		typeIds = append(typeIds, gconv.Uint64(v))
	}
	if projectId <= 0 {
		response.RespJsonExit(req, gcode.CodeValidationFailed.Code(), "所属项目ID必须为正整数")
	}
	if len(typeIds) == 0 {
		response.RespJsonExit(req, gcode.CodeValidationFailed.Code(), "报告类型ID列表不能为空")
	}
	wdkReportAddReq := &v1.WdkReportAddReq{
		ProjectId: projectId,
		TypeIds:   typeIds,
	}
	// 检查新增文档库上传报告记录权限
	_, err = service.WdkReport().AuthAdd(ctx, wdkReportAddReq.ProjectId)
	if err != nil {
		response.RespJsonExit(req, code.AddWdkReportRecordFailed.Code(), code.AddWdkReportRecordFailed.Message()+": "+err.Error())
	}
	// 获取上传报告信息
	file := req.GetUploadFile("upload-report")
	if file == nil {
		response.RespJsonExit(req, code.AddWdkReportRecordFailed.Code(), code.AddWdkReportRecordFailed.Message()+": 获取上传报告信息失败")
	}
	// 上传报告
	var report *upload.FileInfo
	report, err = upload.Upload.UploadFile(file, "wdk/report")
	if err != nil {
		response.RespJsonExit(req, code.AddWdkReportRecordFailed.Code(), code.AddWdkReportRecordFailed.Message()+": "+err.Error())
	}
	// 新增文档库上传报告记录
	err = service.WdkReport().AddWdkReport(ctx, wdkReportAddReq, report)
	if err != nil {
		response.RespJsonExit(req, code.AddWdkReportRecordFailed.Code(), code.AddWdkReportRecordFailed.Message()+": "+err.Error())
	}

	response.Succ(req, v1.WdkReportAddRes{
		TypeIds:  wdkReportAddReq.TypeIds,
		FileInfo: report,
	})
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
