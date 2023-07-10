package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"herbalBody/model"
	"herbalBody/mylog"
	"time"
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
	mylog.Log.Infof("sql: mysql dsn=%s", getMysqlDSN())
	mylog.Log.Info("sql: mysql initializing...")
	var db *sql.DB
	var err error
	success := make(chan struct{}, 1)
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	go func() {
		for {
			db, err = sql.Open("mysql", getMysqlDSN())
			if err != nil || db.Ping() != nil {
				mylog.Log.Error("sql: mysql initialized failed,try again...")
				time.Sleep(2 * time.Second)
				continue
			} else {
				success <- struct{}{}
				mylog.Log.Println("sql: mysql initializes successfully")
				return
			}
		}
	}()
	select {
	case <-timeout.Done():
		panic("sql: mysql initialized timeout!")
	case <-success:
	}
	DB = db

}

func GetDB() *sql.DB {
	return DB
}

func GetGDB() *gorm.DB {
	return GDB
}

// ConnectGorm Connect gorm连接
func ConnectGorm() {
	var db *gorm.DB
	var err error
	success := make(chan struct{}, 1)
	timeout, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	go func() {
		for {
			db, err = gorm.Open(mysql.Open(getMysqlDSN()), &gorm.Config{})
			if err != nil {
				mylog.Log.Error("gorm: mysql initialized failed,try again...")
				time.Sleep(2 * time.Second)
				continue
			} else {
				success <- struct{}{}
				mylog.Log.Println("gorm: mysql initializes successfully")
				return
			}
		}
	}()
	select {
	case <-timeout.Done():
		panic("gorm: mysql initialized timeout!")
	case <-success:
	}

	GDB = db
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
