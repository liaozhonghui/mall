package core

import (
	"fmt"

	"github.com/spf13/viper"
)

var defaultConfig = "./configs/config.yaml"

func InitConfig(configFile string) (err error) {
	if configFile == "" {
		configFile = defaultConfig
	}
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("config read failed:", err)
		return err
	}

	if err = viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Println("config unmarshal failed:", err)
		return err
	}

	return nil
}
