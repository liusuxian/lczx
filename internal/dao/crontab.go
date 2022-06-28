// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"lczx/internal/dao/internal"
)

// internalCrontabDao is internal type for wrapping internal DAO implements.
type internalCrontabDao = *internal.CrontabDao

// crontabDao is the data access object for table crontab.
// You can define custom methods on it to extend its functionality as you wish.
type crontabDao struct {
	internalCrontabDao
}

var (
	// Crontab is globally public accessible object for table crontab operations.
	Crontab = crontabDao{
		internal.NewCrontabDao(),
	}
)

// Fill with you ideas below.