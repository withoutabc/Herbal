package main

import (
	"herbalBody/anotherone/api2"
	"herbalBody/anotherone/conf"
	"herbalBody/anotherone/dao2"
)

func main() {
	conf.ReadConf()
	dao2.InitDB2()
	dao2.ConnectGorm2()
	api2.InitMyRouters()
}
