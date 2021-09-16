package process

import (
	"container/list"
	"github.com/solo-kingdom/nggo.git/pkg/data"
	"github.com/spf13/viper"
)

func Process() {
	dp := viper.GetString("data")
	fm := viper.GetString("format")
	fr := data.NewFileReader(dp)
	rs := list.New()
	for fr.HasNext() {
		line := fr.NextLine()
		record := data.Transfer(line, fm)
		//fmt.Println(record)
		rs.PushBack(record)
		break
	}
}
