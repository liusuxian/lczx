package download

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"lczx/utility/crypto"
)

type fileDownloadAliyunOSS struct {
	Bucket          string
	Endpoint        string
	AccessKeyID     []byte
	AccessKeySecret []byte
}

var FileDownloadAliyunOSS = fileDownloadAliyunOSS{}

func init() {
	ctx := gctx.New()
	// 阿里云OSS下载
	FileDownloadAliyunOSS.Bucket = g.Cfg().MustGet(ctx, "upload.aliyunOSS.bucket").String()
	FileDownloadAliyunOSS.Endpoint = g.Cfg().MustGet(ctx, "upload.aliyunOSS.endpoint").String()
	FileDownloadAliyunOSS.AccessKeyID = g.Cfg().MustGet(ctx, "upload.aliyunOSS.accessKeyID").Bytes()
	FileDownloadAliyunOSS.AccessKeySecret = g.Cfg().MustGet(ctx, "upload.aliyunOSS.accessKeySecret").Bytes()
}

// DownloadAliyunOSS 阿里云OSS下载
func (d fileDownloadAliyunOSS) DownloadAliyunOSS(objectKey, filePath string) (err error) {
	// 解密 accessKeyID
	var accessKeyID []byte
	accessKeyID, err = crypto.AesDecrypt(d.AccessKeyID)
	if err != nil {
		return
	}
	// 解密 accessKeySecret
	var accessKeySecret []byte
	accessKeySecret, err = crypto.AesDecrypt(d.AccessKeySecret)
	if err != nil {
		return
	}
	// 连接OSS
	var client *oss.Client
	client, err = oss.New(d.Endpoint, gstr.TrimAll(string(accessKeyID)), gstr.TrimAll(string(accessKeySecret)))
	if err != nil {
		return
	}
	// 获取存储空间
	var bucket *oss.Bucket
	bucket, err = client.Bucket(d.Bucket)
	if err != nil {
		return
	}
	// 下载文件
	err = bucket.DownloadFile(objectKey, filePath, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	return
}
