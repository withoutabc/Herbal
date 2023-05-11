package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConf() {
	viper.SetConfigName("conf")               // name of config file (without extension)
	viper.SetConfigType("yaml")               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./anotherone/conf/") // optionally look for config in the working directory
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
