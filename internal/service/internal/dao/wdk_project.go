// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/service/internal/dao/internal"
)

// internalWdkProjectDao is internal type for wrapping internal DAO implements.
type internalWdkProjectDao = *internal.WdkProjectDao

// wdkProjectDao is the data access object for table wdk_project.
// You can define custom methods on it to extend its functionality as you wish.
type wdkProjectDao struct {
	internalWdkProjectDao
}

var (
	// WdkProject is globally public accessible object for table wdk_project operations.
	WdkProject = wdkProjectDao{
		internal.NewWdkProjectDao(),
	}
)

// Fill with you ideas below.
