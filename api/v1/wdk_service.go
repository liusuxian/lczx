package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"lczx/internal/model/entity"
)

// WdkServiceRecordReq 文档库服务记录请求参数
type WdkServiceRecordReq struct {
	g.Meta    `path:"/record" tags:"WdkServiceRecord" method:"get" summary:"You first wdk/service/record api"`
	ProjectId uint64 `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
}

// WdkServiceRecordRes 文档库服务记录返回参数
type WdkServiceRecordRes struct {
	List []*WdkServiceInfo `json:"list" dc:"文档库服务记录信息列表"` // 文档库服务记录信息列表
}

// WdkServiceInfo 文档库服务记录信息
type WdkServiceInfo struct {
	Record *entity.WdkServiceRecord  `json:"record" dc:"文档库服务记录"` // 文档库服务记录
	Photo  []*entity.WdkServicePhoto `json:"photo" dc:"文档库服务照片"`  // 文档库服务照片
}

// WdkServiceAddReq 文档库新增服务记录请求参数
type WdkServiceAddReq struct {
	g.Meta        `path:"/add" tags:"WdkServiceAdd" method:"post" summary:"You first wdk/service/add api"`
	ProjectId     uint64              `json:"projectId" v:"required|regex:^[1-9]\\d*$#所属项目ID不能为空|所属项目ID必须为正整数" dc:"所属项目ID"` // 所属项目ID
	ServiceTime   *gtime.Time         `json:"serviceTime" v:"required|date#服务时间不能为空|服务时间不是有效的日期时间" dc:"服务时间"`               // 服务时间
	Remark        string              `json:"remark" v:"max-length:255#备注长度不能超过255" dc:"备注"`                                // 备注
	UploadXchFile *ghttp.UploadFile   `json:"uploadXchFile" dc:"上传行程函文件"`                                                   // 上传行程函文件
	UploadPhoto   *ghttp.UploadFile   `json:"uploadPhoto" dc:"上传照片"`                                                        // 上传照片
	UploadPhotos  []*ghttp.UploadFile `json:"uploadPhotos" dc:"上传照片列表"`                                                     // 上传照片列表
}

// WdkServiceAddRes 文档库新增服务记录返回参数
type WdkServiceAddRes struct {
}
