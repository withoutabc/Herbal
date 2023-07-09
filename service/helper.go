package service

import (
	"gorm.io/gorm"
	"herbalBody/mylog"
)

func ResultErrorRollback(tx *gorm.DB, err error) int {
	mylog.Log.Printf("%v", err)

	if err != nil {
		tx.Rollback()
		return 99
	}
	return 0
}
