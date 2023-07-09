package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
	"herbalBody/mylog"
)

var (
	GDB *gorm.DB
	DB  *sql.DB
)

func getMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=Local&parseTime=true",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.ip"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
}

func InitDB() {

	mylog.Log.Info("mysql initializing...")
	db, err := sql.Open("mysql", getMysqlDSN())
	if err != nil {
		mylog.Log.Errorf("connect mysql err:%v", err)
		return
	}
	DB = db
	if err = db.Ping(); err != nil {
		mylog.Log.Error("mysql fails to initialize")
	}

}

func GetDB() *sql.DB {
	return DB
}

func GetGDB() *gorm.DB {
	return GDB
}

// ConnectGorm Connect gorm连接
func ConnectGorm() {
	db, err := gorm.Open(mysql.Open(getMysqlDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	GDB = db
	mylog.Log.Println("连接成功")
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
