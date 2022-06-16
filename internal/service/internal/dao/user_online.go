// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/service/internal/dao/internal"
)

// internalUserOnlineDao is internal type for wrapping internal DAO implements.
type internalUserOnlineDao = *internal.UserOnlineDao

// userOnlineDao is the data access object for table user_online.
// You can define custom methods on it to extend its functionality as you wish.
type userOnlineDao struct {
	internalUserOnlineDao
}

var (
	// UserOnline is globally public accessible object for table user_online operations.
	UserOnline = userOnlineDao{
		internal.NewUserOnlineDao(),
	}
)

// Fill with you ideas below.
