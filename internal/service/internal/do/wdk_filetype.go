// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WdkFiletype is the golang structure of table wdk_filetype for DAO operations like Where/Data.
type WdkFiletype struct {
	g.Meta `orm:"table:wdk_filetype, do:true"`
	FileId interface{} // 上传文件ID
	TypeId interface{} // 上传文件类型ID 详见wdk_filetype_cfg配置
}
