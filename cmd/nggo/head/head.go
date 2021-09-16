package main

import (
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/config"
	"github.com/solo-kingdom/nggo.git/pkg/data"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("run head")
	config.ReadConfig("configs")
	dp := viper.GetString("data")
	fr := data.NewFileReader(dp)
	fmt.Println(fr.HasNext())
	for i := 0; i < 10 && fr.HasNext(); i++ {
		fmt.Println("go")
		fmt.Println(fr.NextLine())
	}
}