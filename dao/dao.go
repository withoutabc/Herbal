package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
	"log"
)

var (
	GDB *gorm.DB
	DB  *sql.DB
)

func InitDB() {
	db, err := sql.Open("mysql", "root:224488@tcp(127.0.0.1:3306)/herbal?charset=utf8mb4&loc=Local&parseTime=true")
	if err != nil {
		log.Fatalf("connect mysql err:%v", err)
		return
	}
	DB = db
	log.Println(db.Ping())
}

func GetDB() *sql.DB {
	return DB
}

func GetGDB() *gorm.DB {
	return GDB
}

// ConnectGorm Connect gorm连接
func ConnectGorm() {
	dsn := "root:224488@tcp(127.0.0.1:3306)/herbal?charset=utf8mb4&parseTime=True&loc=Local"
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
	GDB.AutoMigrate(&model.BasicInfo{})
	GDB.AutoMigrate(&model.BasicData{})
	GDB.AutoMigrate(&model.MotorData{})
	GDB.AutoMigrate(&model.Conclusion{})
}

func WithTransaction(fn func(db *gorm.DB) error) error {
	tx := GDB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := fn(tx); err != nil {
		return err
	}
	return nil
}
