package config

import "github.com/spf13/viper"

// Config config for mark
type Config interface {
	GetDataPath() string
}

// Impl ConfigImpl config impl
type Impl struct {
	dataPath string
}

func (c Impl) GetDataPath() string {
	return c.dataPath
}

func GenConfig() Config {
	var dp = viper.GetString("data")
	var cfg Config = Impl{dataPath: dp}
	return cfg
}
