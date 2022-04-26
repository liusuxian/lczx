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
	"lczx/utility/utils"
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
	var originFilepath, pdfFilepath string
	originFilepath, pdfFilepath, err = u.uploadAction(file, fType, dirPath)
	if err != nil {
		return
	}

	fileInfo = &FileInfo{
		FileName:      file.Filename,
		FileSize:      file.Size,
		OriginFileUrl: u.getUrl(originFilepath),
		PdfFileUrl:    u.getUrl(pdfFilepath),
		FileType:      file.Header.Get("Content-type"),
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
		var originFilepath, pdfFilepath string
		originFilepath, pdfFilepath, err = u.uploadAction(file, fType, dirPath)
		if err != nil {
			return
		}

		fileInfo := &FileInfo{
			FileName:      file.Filename,
			FileSize:      file.Size,
			OriginFileUrl: u.getUrl(originFilepath),
			PdfFileUrl:    u.getUrl(pdfFilepath),
			FileType:      file.Header.Get("Content-type"),
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
func (u FileUploadOSSAdapter) uploadAction(file *ghttp.UploadFile, fType string, dirPath string) (originFilepath, pdfFilepath string, err error) {
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
	// 打开文件
	var fd io.ReadCloser
	fd, err = file.Open()
	if err != nil {
		return
	}
	defer fd.Close()
	// 文件名处理
	filename := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(10))
	fileFullname := filename + gfile.Ext(file.Filename)
	originFilepath = dirPath + "/" + fileFullname
	// 上传原始文件
	err = bucket.PutObject(originFilepath, fd)
	if err != nil {
		return
	}
	// 原始文件转pdf文件
	extName := gfile.ExtName(file.Filename)
	if fType == "file" && extName == "pdf" {
		pdfFilepath = dirPath + "/" + filename + ".pdf"
		return
	}
	if fType == "file" {
		// 下载原始文件
		downloadFilepath := "./cache/download/" + fileFullname
		err = bucket.DownloadFile(originFilepath, downloadFilepath, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
		if err != nil {
			return
		}
		// 转pdf
		var resultPath string
		resultPath, err = utils.ConvertToPDF(downloadFilepath)
		if err != nil {
			// 删除下载的原始文件
			_ = gfile.Remove(downloadFilepath)
			return
		}
		// 上传pdf文件
		if resultPath != "" {
			tmpPdfFilepath := dirPath + "/" + filename + ".pdf"
			// 上传文件
			err = bucket.PutObjectFromFile(tmpPdfFilepath, resultPath)
			if err != nil {
				// 删除下载的原始文件
				_ = gfile.Remove(downloadFilepath)
				// 删除转换的pdf文件
				_ = gfile.Remove(resultPath)
				return
			}
			// 删除下载的原始文件
			_ = gfile.Remove(downloadFilepath)
			// 删除转换的pdf文件
			_ = gfile.Remove(resultPath)
			pdfFilepath = tmpPdfFilepath
			return
		}
	}
	return
}
