package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

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
