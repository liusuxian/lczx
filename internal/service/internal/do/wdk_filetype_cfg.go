// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WdkFiletypeCfg is the golang structure of table wdk_filetype_cfg for DAO operations like Where/Data.
type WdkFiletypeCfg struct {
	g.Meta   `orm:"table:wdk_filetype_cfg, do:true"`
	Id       interface{} // 上传文件类型ID
	Name     interface{} // 类型名称
	Multiple interface{} // 是否同时存在多个文件 0:否 1:是
	Audit    interface{} // 是否需要审核 0:不需要 1:需要
	Step     interface{} // 所属项目阶段 0:未开始 1:合同签约 2:项目启动会 3:服务中 4:合同结束 5:复盘
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 更新时间
}
