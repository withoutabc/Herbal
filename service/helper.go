package service

import (
	"gorm.io/gorm"
	"log"
)

func ResultErrorRollback(tx *gorm.DB, err error) int {
	log.Printf("%v", err)

	if err != nil {
		tx.Rollback()
		return 99
	}
	return 0
}
