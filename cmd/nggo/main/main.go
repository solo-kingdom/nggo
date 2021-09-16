package main

import (
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/config"
	"github.com/solo-kingdom/nggo.git/pkg/process"
	"github.com/spf13/viper"
)

func main() {
	config.ReadConfig("configs")
	dt := viper.Get("data")
	fm := viper.Get("format")
	fmt.Printf("run prcess. [data='%s', format='%s']\n\n", dt, fm)
	process.Process()
}
