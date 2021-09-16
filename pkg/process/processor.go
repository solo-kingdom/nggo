package process

import (
	"github.com/solo-kingdom/nggo.git/pkg/analyst"
	"github.com/solo-kingdom/nggo.git/pkg/data"
	"github.com/spf13/viper"
	"log"
)

func Process() {
	dp := viper.GetString("data")
	fm := viper.GetString("format")
	fr := data.NewFileReader(dp)
	rs := make([]map[string]interface{}, 0, 1000000)
	// convert format
	ffm := ConvertFormat(fm)
	// read log file line by line
	for fr.HasNext() {
		line := fr.NextLine()
		// convert one log line into map structure
		record, err := data.Transfer(&line, ffm)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(record)
		rs = append(rs, record)
	}

	log.Print("total records: ", len(rs))

	// analyst
	ua := analyst.GetUserAnalyst()
	ua.Process(rs)
	ua.Show()
}
