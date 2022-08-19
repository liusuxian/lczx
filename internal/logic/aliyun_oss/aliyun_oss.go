package aliyun_oss

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"lczx/internal/service"
	"lczx/utility/logger"
)

// 阿里云OSS信息
type sAliyunOSS struct {
	bucket             string // bucket名称
	endpointAccelerate string // 传输加速节点
	endpointInternal   string // 内网访问节点
	endpointAccess     string // 外网访问节点
	accessKeyID        []byte // accessKeyID
	accessKeySecret    []byte // accessKeySecret
}

// 上传配置
type uploadCfg struct {
	env      int      // 环境
	imgList  []string // 允许上传的图片类型
	fileList []string // 允许上传的文件类型
	maxSize  int64    // 最大文件上传大小(MB)
}

var (
	upCfg *uploadCfg
)

func init() {
	service.RegisterAliyunOSS(newAliyunOSS())
}

// 阿里云OSS服务
func newAliyunOSS() *sAliyunOSS {
	ctx := gctx.New()
	upCfg = &uploadCfg{
		env:      g.Cfg().MustGet(ctx, "aliyunoss.env", 0).Int(),
		imgList:  g.Cfg().MustGet(ctx, "aliyunoss.upload.imgList", []string{}).Strings(),
		fileList: g.Cfg().MustGet(ctx, "aliyunoss.upload.fileList", []string{}).Strings(),
		maxSize:  g.Cfg().MustGet(ctx, "aliyunoss.upload.maxSize", 5120).Int64(),
	}

	var oss *sAliyunOSS
	if upCfg.env == 0 {
		// 开发环境
		oss = &sAliyunOSS{
			bucket:             g.Cfg().MustGet(ctx, "aliyunoss.dev.bucket").String(),
			endpointAccelerate: g.Cfg().MustGet(ctx, "aliyunoss.dev.endpointAccelerate").String(),
			endpointInternal:   g.Cfg().MustGet(ctx, "aliyunoss.dev.endpointInternal").String(),
			endpointAccess:     g.Cfg().MustGet(ctx, "aliyunoss.dev.endpointAccess").String(),
			accessKeyID:        g.Cfg().MustGet(ctx, "aliyunoss.dev.accessKeyID").Bytes(),
			accessKeySecret:    g.Cfg().MustGet(ctx, "aliyunoss.dev.accessKeySecret").Bytes(),
		}
	} else if upCfg.env == 1 {
		// 正式环境
		oss = &sAliyunOSS{
			bucket:             g.Cfg().MustGet(ctx, "aliyunoss.prod.bucket").String(),
			endpointAccelerate: g.Cfg().MustGet(ctx, "aliyunoss.prod.endpointAccelerate").String(),
			endpointInternal:   g.Cfg().MustGet(ctx, "aliyunoss.prod.endpointInternal").String(),
			endpointAccess:     g.Cfg().MustGet(ctx, "aliyunoss.prod.endpointAccess").String(),
			accessKeyID:        g.Cfg().MustGet(ctx, "aliyunoss.prod.accessKeyID").Bytes(),
			accessKeySecret:    g.Cfg().MustGet(ctx, "aliyunoss.prod.accessKeySecret").Bytes(),
		}
	} else {
		logger.Panic(ctx, "Init AliyunOSS Config Error")
	}

	return oss
}
