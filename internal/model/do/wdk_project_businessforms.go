// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WdkProjectBusinessforms is the golang structure of table wdk_project_businessforms for DAO operations like Where/Data.
type WdkProjectBusinessforms struct {
	g.Meta        `orm:"table:wdk_project_businessforms, do:true"`
	ProjectId     interface{} // 文档库项目ID
	BusinessForms interface{} // 业态 0:住宅 1:小高层 2:高层 3:超高层 4:公寓 5:合院 6:叠墅 7:排屋 8:多层 9:会所 10:商住 11:综合体 12:产业园 13:酒店 14:酒店式公寓 15:商业 16:普通商业 17:公共配套 18:办公 19:公寓式办公 20:厂房
}
