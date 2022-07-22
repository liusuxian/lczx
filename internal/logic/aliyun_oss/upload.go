package aliyun_oss

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"lczx/internal/consts"
	"lczx/internal/model"
	"lczx/internal/service"
	"lczx/utility/crypto"
	"lczx/utility/utils"
	"strconv"
	"strings"
	"time"
)

// UploadImg 单图片上传
func (s *sAliyunOSS) UploadImg(file *ghttp.UploadFile, dirPath string) (fileInfo *model.UploadFileInfo, err error) {
	return s.uploadByType(file, dirPath, "img")
}

// UploadFile 单文件上传
func (s *sAliyunOSS) UploadFile(file *ghttp.UploadFile, dirPath string) (fileInfo *model.UploadFileInfo, err error) {
	return s.uploadByType(file, dirPath, "file")
}

// UploadImgs 多图片上传
func (s *sAliyunOSS) UploadImgs(files []*ghttp.UploadFile, dirPath string) (fileInfos []*model.UploadFileInfo, err error) {
	return s.uploadBathByType(files, dirPath, "img")
}

// UploadFiles 多文件上传
func (s *sAliyunOSS) UploadFiles(files []*ghttp.UploadFile, dirPath string) (fileInfos []*model.UploadFileInfo, err error) {
	return s.uploadBathByType(files, dirPath, "file")
}

// GetAccessUrl 文件访问Url
func (s *sAliyunOSS) GetAccessUrl(ctx context.Context, filePath string) (fileUrl string, err error) {
	// 从缓存获取url
	strList := gstr.Split(filePath, "/")
	cacheKey := utils.RedisKey(consts.CacheAccessUrlPrefix, strList[len(strList)-1])
	var cacheVal *gvar.Var
	cacheVal = service.Cache().GetCache(ctx, cacheKey)
	if cacheVal != nil {
		fileUrl = gconv.String(cacheVal)
		if fileUrl != "" {
			return
		}
	}
	// 解密 accessKeyID
	var accessKeyID []byte
	accessKeyID, err = crypto.AesDecrypt(s.accessKeyID)
	if err != nil {
		return
	}
	// 解密 accessKeySecret
	var accessKeySecret []byte
	accessKeySecret, err = crypto.AesDecrypt(s.accessKeySecret)
	if err != nil {
		return
	}
	// 连接OSS
	var client *oss.Client
	client, err = oss.New(s.endpointAccess, gstr.TrimAll(string(accessKeyID)), gstr.TrimAll(string(accessKeySecret)))
	if err != nil {
		return
	}
	// 获取存储空间
	var bucket *oss.Bucket
	bucket, err = client.Bucket(s.bucket)
	if err != nil {
		return
	}
	// 授权访问
	fileUrl, err = bucket.SignURL(filePath, oss.HTTPGet, 3610)
	if err != nil {
		return
	}
	// 添加缓存
	err = service.Cache().SetCache(ctx, cacheKey, fileUrl, time.Hour)
	return
}

// 文件上传 img|file
func (s *sAliyunOSS) uploadByType(file *ghttp.UploadFile, dirPath string, fType string) (fileInfo *model.UploadFileInfo, err error) {
	if file == nil {
		err = gerror.New("未上传任何文件")
		return
	}
	// 检测文件类型
	var typeList []string
	if fType == "img" {
		typeList = upCfg.imgList
	} else if fType == "file" {
		typeList = upCfg.fileList
	}
	// 判断上传文件类型是否合法
	if !s.checkFileType(typeList, file.Filename) {
		err = gerror.New("上传文件类型错误，只能包含后缀为 " + gstr.Join(typeList, ",") + " 的文件。")
		return
	}
	// 检查上传文件大小是否合法
	if !s.checkSize(upCfg.maxSize, file.Size) {
		err = gerror.New("上传文件超过最大尺寸：" + gconv.String(upCfg.maxSize) + "MB")
		return
	}
	// 执行上传
	var originFileUrl, pdfFileUrl string
	originFileUrl, pdfFileUrl, err = s.uploadAction(file, fType, dirPath)
	if err != nil {
		return
	}

	fileInfo = &model.UploadFileInfo{
		FileName:      file.Filename,
		FileSize:      file.Size,
		OriginFileUrl: originFileUrl,
		PdfFileUrl:    pdfFileUrl,
		FileType:      file.Header.Get("Content-type"),
	}
	return
}

// 批量上传 img|file
func (s *sAliyunOSS) uploadBathByType(files []*ghttp.UploadFile, dirPath string, fType string) (fileInfos []*model.UploadFileInfo, err error) {
	if len(files) == 0 {
		err = gerror.New("未上传任何文件")
		return
	}
	// 检测文件类型
	var typeList []string
	if fType == "img" {
		typeList = upCfg.imgList
	} else if fType == "file" {
		typeList = upCfg.fileList
	}
	// 循环检测
	for _, file := range files {
		// 判断上传文件类型是否合法
		if !s.checkFileType(typeList, file.Filename) {
			err = gerror.New("上传文件类型错误，只能包含后缀为 " + gstr.Join(typeList, ",") + " 的文件。")
			return
		}
		// 检查上传文件大小是否合法
		if !s.checkSize(upCfg.maxSize, file.Size) {
			err = gerror.New("上传文件超过最大尺寸：" + gconv.String(upCfg.maxSize) + "MB")
			return
		}
	}
	// 循环执行上传
	fileInfos = make([]*model.UploadFileInfo, 0, len(files))
	for _, file := range files {
		var originFileUrl, pdfFileUrl string
		originFileUrl, pdfFileUrl, err = s.uploadAction(file, fType, dirPath)
		if err != nil {
			return
		}

		fileInfo := &model.UploadFileInfo{
			FileName:      file.Filename,
			FileSize:      file.Size,
			OriginFileUrl: originFileUrl,
			PdfFileUrl:    pdfFileUrl,
			FileType:      file.Header.Get("Content-type"),
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return
}

// 判断上传文件类型是否合法
func (s *sAliyunOSS) checkFileType(typeList []string, filename string) bool {
	fileType := gfile.ExtName(filename)
	return gstr.InArray(typeList, fileType)
}

// 检查上传文件大小是否合法
func (s *sAliyunOSS) checkSize(configSize int64, fileSize int64) bool {
	return configSize*1024*1024 >= fileSize
}

// 上传
func (s *sAliyunOSS) uploadAction(file *ghttp.UploadFile, fType string, dirPath string) (originFileUrl, pdfFileUrl string, err error) {
	// 图片/文件名处理
	filename := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(10))
	fileFullname := filename + gfile.Ext(file.Filename)
	// 文件处理
	var resultPath string
	extName := gfile.ExtName(file.Filename)
	localFilepath := "./cache/local/" + fileFullname
	pdfFilepath := dirPath + "/" + filename + ".pdf"
	if fType == "file" {
		pdfFileUrl = pdfFilepath
		if extName != "pdf" {
			// 保存本地临时文件
			var localFilename string
			localFilename, err = file.Save("./cache/local/")
			if err != nil {
				return
			}
			// 本地临时文件改名
			err = gfile.Rename("./cache/local/"+localFilename, localFilepath)
			if err != nil {
				// 删除本地临时文件
				_ = gfile.Remove("./cache/local/" + localFilename)
				return
			}
			// 原始文件转pdf文件
			resultPath, err = utils.ConvertToPDF(localFilepath, "./cache/pdf/")
			if err != nil {
				// 删除本地临时文件
				_ = gfile.Remove(localFilepath)
				return
			}
		}
	}
	// 解密 accessKeyID
	var accessKeyID []byte
	accessKeyID, err = crypto.AesDecrypt(s.accessKeyID)
	if err != nil {
		return
	}
	// 解密 accessKeySecret
	var accessKeySecret []byte
	accessKeySecret, err = crypto.AesDecrypt(s.accessKeySecret)
	if err != nil {
		return
	}
	// 连接OSS
	var client *oss.Client
	client, err = oss.New(s.endpointAccelerate, gstr.TrimAll(string(accessKeyID)), gstr.TrimAll(string(accessKeySecret)))
	if err != nil {
		return
	}
	// 获取存储空间
	var bucket *oss.Bucket
	bucket, err = client.Bucket(s.bucket)
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
	// 上传原始文件
	originFilepath := dirPath + "/" + fileFullname
	err = bucket.PutObject(originFilepath, fd)
	if err != nil {
		// 删除本地临时文件
		_ = gfile.Remove(localFilepath)
		// 删除转换的pdf文件
		_ = gfile.Remove(resultPath)
		return
	}
	originFileUrl = originFilepath
	// 上传pdf文件
	if fType == "file" && extName != "pdf" && resultPath != "" {
		// 上传文件
		err = bucket.PutObjectFromFile(pdfFilepath, resultPath)
		// 删除本地临时文件
		_ = gfile.Remove(localFilepath)
		// 删除转换的pdf文件
		_ = gfile.Remove(resultPath)
	}
	return
}
