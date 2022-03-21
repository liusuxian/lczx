// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/service/internal/dao/internal"
)

// deptDao is the data access object for table dept.
// You can define custom methods on it to extend its functionality as you wish.
type deptDao struct {
	*internal.DeptDao
}

var (
	// Dept is globally public accessible object for table dept operations.
	Dept = deptDao{
		internal.NewDeptDao(),
	}
)

// Fill with you ideas below.