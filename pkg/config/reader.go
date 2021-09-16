package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(pt string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pt)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error occurred while read config : %s", err))
	}
}
