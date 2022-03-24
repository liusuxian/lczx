package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DeptInfo 部门信息
type DeptInfo struct {
	Id   uint   `json:"id"   dc:"部门ID"` // 部门ID
	Name string `json:"name" dc:"部门名称"` // 部门名称
}

type DeptListReq struct {
	g.Meta `path:"/list" tags:"DeptList" method:"post" summary:"You first /dept/list api"`
}
type DeptListRes struct {
	List []*DeptInfo `json:"list" dc:"部门列表"` // 部门列表
}

type DeptAddReq struct {
	g.Meta `path:"/add" tags:"DeptAdd" method:"post" summary:"You first /dept/add api"`
	Name   string `json:"name" v:"required|max-length:10|regex:^[\u4e00-\u9fa5]+$#部门名称不能为空|部门名称不能超过10个字|部门名称必须为纯中文" dc:"部门名称"` // 部门名称
}
type DeptAddRes struct {
	Dept *DeptInfo `json:"dept" dc:"部门信息"` // 部门信息
}

type DeptDeleteReq struct {
	g.Meta `path:"/delete" tags:"DeptDelete" method:"post" summary:"You first /dept/delete api"`
	Id     uint `json:"id" v:"required|regex:^[1-9][0-9]*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"` // 部门ID
}
type DeptDeleteRes struct {
	Id uint `json:"id" dc:"部门ID"` // 部门ID
}

type DeptUpdateReq struct {
	g.Meta `path:"/update" tags:"DeptUpdate" method:"post" summary:"You first /dept/update api"`
	Id     uint   `json:"id" v:"required|regex:^[1-9][0-9]*$#部门ID不能为空|部门ID必须为正整数" dc:"部门ID"`                                   // 部门ID
	Name   string `json:"name" v:"required|max-length:10|regex:^[\u4e00-\u9fa5]+$#部门名称不能为空|部门名称不能超过10个字|部门名称必须为纯中文" dc:"部门名称"` // 部门名称
}
type DeptUpdateRes struct {
	Dept *DeptInfo `json:"dept" dc:"部门信息"` // 部门信息
}
