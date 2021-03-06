// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/dao/internal"
)

// internalWdkServicePhotoDao is internal type for wrapping internal DAO implements.
type internalWdkServicePhotoDao = *internal.WdkServicePhotoDao

// wdkServicePhotoDao is the data access object for table wdk_service_photo.
// You can define custom methods on it to extend its functionality as you wish.
type wdkServicePhotoDao struct {
	internalWdkServicePhotoDao
}

var (
	// WdkServicePhoto is globally public accessible object for table wdk_service_photo operations.
	WdkServicePhoto = wdkServicePhotoDao{
		internal.NewWdkServicePhotoDao(),
	}
)

// Fill with you ideas below.
