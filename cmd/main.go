package main

import (
	"fmt"
	"github.com/spf13/viper"
	"herbalBody/api"
	"herbalBody/dao"
	"herbalBody/mylog"
)

func main() {
	readConf()
	mylog.InitLog()
	dao.InitDB()
	dao.ConnectGorm()
	api.InitRouter()
}
func readConf() {
	fmt.Println("reading config...")
	viper.SetConfigName("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./etc/")      // optionally look for config in the working directory
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	} else {
		fmt.Println("config read success")
	}
}
