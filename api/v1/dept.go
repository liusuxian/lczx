package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"lczx/internal/model/entity"
)

type DeptListReq struct {
	g.Meta `path:"/getDeptList" tags:"GetDeptList" method:"post" summary:"You first getDeptList api"`
}
type DeptListRes struct {
	DeptList []*entity.Dept `json:"deptList" dc:"部门列表"` // 部门列表
}
