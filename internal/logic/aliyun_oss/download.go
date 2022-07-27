package aliyun_oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/text/gstr"
	"lczx/utility/crypto"
)

// AuthorizationDownload 授权给第三方下载
func (s *sAliyunOSS) AuthorizationDownload(filePath string) (fileUrl string, err error) {
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
	// 授权访问
	fileUrl, err = bucket.SignURL(filePath, oss.HTTPGet, 60)
	return
}
