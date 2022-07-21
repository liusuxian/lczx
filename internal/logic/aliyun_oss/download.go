package aliyun_oss

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"io"
	"lczx/utility/crypto"
	"net/url"
)

// Download 文件下载
func (s *sAliyunOSS) Download(ctx context.Context, filePath string) (body io.ReadCloser, err error) {
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
	// 使用签名URL将OSS文件下载到流
	var signedURL string
	signedURL, err = bucket.SignURL(filePath, oss.HTTPGet, 60)
	if err != nil {
		return
	}
	body, err = bucket.GetObjectWithURL(signedURL)
	if err != nil {
		return
	}
	// 获取文件名
	strList := gstr.Split(filePath, "/")
	filename := strList[len(strList)-1]
	// 设置响应头
	request := g.RequestFromCtx(ctx)
	request.Response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename=%s`, url.QueryEscape(filename)))
	request.Response.Header().Set("Access-Control-Expose-Headers", "content-disposition")
	return
}
