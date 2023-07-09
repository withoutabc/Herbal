package dao2

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
)

var (
	GDB *gorm.DB
)

// ConnectGorm2 Connect gorm连接
func ConnectGorm2() {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("gorm: failed to connect database")
	}
	GDB = db
	log.Info("gorm init success")
	GDB.AutoMigrate(&model.Questions{})
	GDB.AutoMigrate(&model.Option{})
	GDB.AutoMigrate(&model.User{})
	GDB.AutoMigrate(&model.Questionnaires{})
	GDB.AutoMigrate(&model.Version{})
	GDB.AutoMigrate(&model.Title{})
	GDB.AutoMigrate(&model.List{})
	GDB.AutoMigrate(&model.Grade{})
	GDB.AutoMigrate(&model.Signature{})
}
