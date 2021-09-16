package main

import (
	"flag"
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/config"
	"github.com/solo-kingdom/nggo.git/pkg/process"
	"github.com/spf13/viper"
	"os"
)

func main() {
	config.ReadConfig("configs")
	data := flag.String("data", "", "nginx log file path")
	flag.Parse()

	if len(*data) > 0 {
		viper.Set("data", *data)
	}

	// print information
	dt := viper.Get("data")
	fm := viper.Get("format")
	fmt.Printf("run prcess. [args='%v', data='%v', format='%v']\n\n", os.Args, dt, fm)
	process.Process()
}
