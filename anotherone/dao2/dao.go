package dao2

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
	"log"
)

var (
	GDB *gorm.DB
	DB  *sql.DB
)

func InitDB2() {
	db, err := sql.Open("mysql", viper.GetString("mysql.dsn"))
	if err != nil {
		log.Fatalf("connect mysql err:%v", err)
		return
	}
	DB = db
	log.Println(db.Ping())
}

func GetDB2() *sql.DB {
	return DB
}

func GetGDB2() *gorm.DB {
	return GDB
}

// ConnectGorm Connect gorm连接
func ConnectGorm2() {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	GDB = db
	log.Println("连接成功")
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
