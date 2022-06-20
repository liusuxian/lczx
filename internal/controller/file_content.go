package controller

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"golang.org/x/net/context"
	"io"
	v1 "lczx/api/v1"
	"lczx/internal/code"
	"lczx/internal/download"
	"os"
)

var (
	FileContent = cFileContent{}
)

type cFileContent struct{}

// GetFileContent 获取文件内容
func (c *cFileContent) GetFileContent(ctx context.Context, req *v1.FileContentGetReq) (res *v1.FileContentGetRes, err error) {
	// 下载文件
	list := gstr.SplitAndTrim(req.FileUrl, "/")
	filePath := "./cache/local/" + list[len(list)-1]
	if !gfile.Exists(filePath) {
		// 文件不存在则下载
		objectKey := gstr.Join(list[2:], "/")
		err = download.AliyunOSS.Download(objectKey, filePath)
		if err != nil {
			err = gerror.WrapCode(code.GetFileContentFailed, err)
			return
		}
	}
	// 读取文件内容
	var file *os.File
	file, err = gfile.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		err = gerror.WrapCode(code.GetFileContentFailed, err)
		return
	}
	defer file.Close()
	_, err = io.Copy(g.RequestFromCtx(ctx).Response.Writer, file)
	return
}
