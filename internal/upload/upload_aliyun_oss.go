package upload

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"lczx/utility/crypto"
	"strconv"
	"strings"
)

func init() {
	var adp FileUploadAdapter
	if upType == 1 {
		// 使用阿里云OSS上传
		adp = FileUploadOSSAdapter{
			Bucket:          g.Cfg().MustGet(ctx, "upload.aliyunOSS.bucket").String(),
			Endpoint:        g.Cfg().MustGet(ctx, "upload.aliyunOSS.endpoint").String(),
			RawUrl:          g.Cfg().MustGet(ctx, "upload.aliyunOSS.rawUrl").String(),
			AccessKeyID:     g.Cfg().MustGet(ctx, "upload.aliyunOSS.accessKeyID").Bytes(),
			AccessKeySecret: g.Cfg().MustGet(ctx, "upload.aliyunOSS.accessKeySecret").Bytes(),
		}
		Upload = &upload{
			upAdapter: adp,
		}
	}
}

type FileUploadOSSAdapter struct {
	Bucket          string
	Endpoint        string
	RawUrl          string
	AccessKeyID     []byte
	AccessKeySecret []byte
}

func (u FileUploadOSSAdapter) UploadImg(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error) {
	return u.uploadByType(file, dirPath, "img")
}

func (u FileUploadOSSAdapter) UploadFile(file *ghttp.UploadFile, dirPath string) (fileInfo *FileInfo, err error) {
	return u.uploadByType(file, dirPath, "file")
}

func (u FileUploadOSSAdapter) UploadImgs(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error) {
	return u.uploadBathByType(files, dirPath, "img")
}

func (u FileUploadOSSAdapter) UploadFiles(files []*ghttp.UploadFile, dirPath string) (fileInfos []*FileInfo, err error) {
	return u.uploadBathByType(files, dirPath, "file")
}

// 文件上传 img|file
func (u FileUploadOSSAdapter) uploadByType(file *ghttp.UploadFile, dirPath string, fType string) (fileInfo *FileInfo, err error) {
	if file == nil {
		err = gerror.New("未上传任何文件")
		return
	}
	// 检测文件类型
	var typeList []string
	if fType == "img" {
		typeList = imgList
	} else if fType == "file" {
		typeList = fileList
	}
	// 判断上传文件类型是否合法
	if !u.checkFileType(typeList, file.Filename) {
		err = gerror.New("上传文件类型错误，只能包含后缀为 " + gstr.Join(typeList, ",") + " 的文件。")
		return
	}
	// 检查上传文件大小是否合法
	if !u.checkSize(maxSize, file.Size) {
		err = gerror.New("上传文件超过最大尺寸：" + gconv.String(maxSize) + "MB")
		return
	}
	// 执行上传
	var filepath string
	filepath, err = u.uploadAction(file, dirPath)
	if err != nil {
		return
	}

	fileInfo = &FileInfo{
		FileName: file.Filename,
		FileSize: file.Size,
		FileUrl:  u.getUrl(filepath),
		FileType: file.Header.Get("Content-type"),
	}
	return
}

// 批量上传 img|file
func (u FileUploadOSSAdapter) uploadBathByType(files []*ghttp.UploadFile, dirPath string, fType string) (fileInfos []*FileInfo, err error) {
	if len(files) == 0 {
		err = gerror.New("未上传任何文件")
		return
	}
	// 检测文件类型
	var typeList []string
	if fType == "img" {
		typeList = imgList
	} else if fType == "file" {
		typeList = fileList
	}
	// 循环检测
	for _, file := range files {
		// 判断上传文件类型是否合法
		if !u.checkFileType(typeList, file.Filename) {
			err = gerror.New("上传文件类型错误，只能包含后缀为 " + gstr.Join(typeList, ",") + " 的文件。")
			return
		}
		// 检查上传文件大小是否合法
		if !u.checkSize(maxSize, file.Size) {
			err = gerror.New("上传文件超过最大尺寸：" + gconv.String(maxSize) + "MB")
			return
		}
	}
	// 循环执行上传
	fileInfos = make([]*FileInfo, 0, len(files))
	for _, file := range files {
		var filepath string
		filepath, err = u.uploadAction(file, dirPath)
		if err != nil {
			return
		}

		fileInfo := &FileInfo{
			FileName: file.Filename,
			FileSize: file.Size,
			FileUrl:  u.getUrl(filepath),
			FileType: file.Header.Get("Content-type"),
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return
}

// 判断上传文件类型是否合法
func (u FileUploadOSSAdapter) checkFileType(typeList []string, filename string) bool {
	fileType := gfile.ExtName(filename)
	return gstr.InArray(typeList, fileType)
}

// 检查上传文件大小是否合法
func (u FileUploadOSSAdapter) checkSize(configSize int64, fileSize int64) bool {
	return configSize*1024*1024 >= fileSize
}

// 获取上传文件的Url
func (u FileUploadOSSAdapter) getUrl(filepath string) string {
	return u.RawUrl + "/" + filepath
}

// 上传到阿里云OSS操作
func (u FileUploadOSSAdapter) uploadAction(file *ghttp.UploadFile, dirPath string) (filepath string, err error) {
	// 文件名处理
	filename := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(10))
	filename = filename + gfile.Ext(file.Filename)
	filepath = dirPath + "/" + filename
	// 解密 accessKeyID
	var accessKeyID []byte
	accessKeyID, err = crypto.AesDecrypt(u.AccessKeyID)
	if err != nil {
		return
	}
	// 解密 accessKeySecret
	var accessKeySecret []byte
	accessKeySecret, err = crypto.AesDecrypt(u.AccessKeySecret)
	if err != nil {
		return
	}
	// 连接OSS
	var client *oss.Client
	client, err = oss.New(u.Endpoint, gstr.TrimAll(string(accessKeyID)), gstr.TrimAll(string(accessKeySecret)))
	if err != nil {
		return
	}
	// 获取存储空间
	var bucket *oss.Bucket
	bucket, err = client.Bucket(u.Bucket)
	if err != nil {
		return
	}
	// 设置Option
	options := []oss.Option{
		// 指定该Object被下载时网页的缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		oss.ContentDisposition("attachment;filename=" + file.Filename),
		// 指定该Object被下载时的内容编码格式。
		// oss.ContentEncoding("UTF-8"),
		// 指定过期时间。
		// oss.Expires(time.Date(2023, time.January, 30, 23, 0, 0, 0, time.UTC)),
		// 指定对返回的Key进行编码，目前支持URL编码。
		oss.EncodingType("url"),
		// 指定Object的存储类型。
		oss.ObjectStorageClass(oss.StorageStandard),
		// 指定Object的访问权限。
		// oss.ObjectACL("private"),
		// 指定CopyObject操作时是否覆盖同名目标Object。此处设置为true，表示禁止覆盖同名Object。
		oss.ForbidOverWrite(true),
	}
	// 打开文件
	var fd io.ReadCloser
	fd, err = file.Open()
	if err != nil {
		return
	}
	defer fd.Close()
	// 上传文件
	err = bucket.PutObject(filepath, fd, options...)
	return
}
