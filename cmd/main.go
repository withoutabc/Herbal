package main

import (
	"herbalBody/api"
	"herbalBody/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
