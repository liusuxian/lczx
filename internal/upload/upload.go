package upload

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// FileInfo 上传的文件信息
type FileInfo struct {
	FileName      string `json:"fileName" dc:"文件名"`          // 文件名
	FileSize      int64  `json:"fileSize" dc:"文件大小"`         // 文件大小
	OriginFileUrl string `json:"originFileUrl" dc:"原始文件Url"` // 原始文件Url
	PdfFileUrl    string `json:"pdfFileUrl" dc:"pdf文件Url"`   // pdf文件Url
	FileType      string `json:"fileType" dc:"文件类型"`         // 文件类型
}

// FileUploadAdapter 文件上传适配器接口
type FileUploadAdapter interface {
	UploadImg(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error)
	UploadFile(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error)
	UploadImgs(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error)
	UploadFiles(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error)
	GetUrl(filepath string) string
}

type upload struct {
	upAdapter FileUploadAdapter
}

var (
	ctx      = gctx.New()
	upType   = g.Cfg().MustGet(ctx, "upload.type", 1).Int()
	imgList  = g.Cfg().MustGet(ctx, "upload.imgList", []string{}).Strings()
	fileList = g.Cfg().MustGet(ctx, "upload.fileList", []string{}).Strings()
	maxSize  = g.Cfg().MustGet(ctx, "upload.maxSize", 5120).Int64()
	Upload   *upload
)

// UploadImg 单图片上传
func (u upload) UploadImg(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error) {
	return u.upAdapter.UploadImg(file, dirPath)
}

// UploadFile 单文件上传
func (u upload) UploadFile(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error) {
	return u.upAdapter.UploadFile(file, dirPath)
}

// UploadImgs 多图片上传
func (u upload) UploadImgs(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error) {
	return u.upAdapter.UploadImgs(files, dirPath)
}

// UploadFiles 多文件上传
func (u upload) UploadFiles(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error) {
	return u.upAdapter.UploadFiles(files, dirPath)
}

// GetUrl 获取上传文件的Url
func (u upload) GetUrl(filepath string) string {
	return u.upAdapter.GetUrl(filepath)
}
