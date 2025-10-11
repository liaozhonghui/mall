package core

import (
	"fmt"

	"github.com/spf13/viper"
)

var defaultConfig = "../../configs/config.yaml"

func InitConfig(configPath string) error {
	if configPath == "" {
		configPath = defaultConfig
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v", err)
		return err
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		fmt.Printf("Unable to decode into struct: %v", err)
		return err
	}
	return nil
}
