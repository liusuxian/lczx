package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/service"
)

var (
	WdkFile = cWdkFile{}
)

type cWdkFile struct{}

// Record 文档库上传文件记录
func (c *cWdkFile) Record(ctx context.Context, req *v1.WdkFileRecordReq) (res *v1.WdkFileRecordRes, err error) {
	var list []*v1.WdkFileRecordInfo
	list, err = service.WdkFile().GetWdkFileRecord(ctx, req.ProjectId)
	if err != nil {
		err = gerror.WrapCode(code.GetWdkFileRecordFailed, err)
		return
	}

	res = &v1.WdkFileRecordRes{List: list}
	return
}
