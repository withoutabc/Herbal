package dao2

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log2 "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
)

var (
	GDB *gorm.DB
	DB  *sql.DB
)

func InitDB2() {
	db, err := sql.Open("mysql", viper.GetString("mysql.dsn"))
	if err != nil {
		log2.Error("mysql: connect failed..., err: :", err)
		return
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log2.Error(err)
		return
	}
	log2.Info("mysql init success")

}

func GetDB2() *sql.DB {
	return DB
}

func GetGDB2() *gorm.DB {
	return GDB
}

// ConnectGorm2 Connect gorm连接
func ConnectGorm2() {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log2.Error("gorm: failed to connect database")
	}
	GDB = db
	log2.Info("gorm init success")
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
