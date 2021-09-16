package process

import (
	"encoding/json"
	"github.com/solo-kingdom/nggo.git/pkg/convertor"
	"github.com/solo-kingdom/nggo.git/pkg/helper"
	"log"
	"strconv"
	"strings"
	"time"
)

//goland:noinspection GoUnusedGlobalVariable
var FormatMapper = map[string]convertor.Convertor{
	"$remote_user":  {Run: StringFormat, Name: "user"},
	"$request_body": {Run: JsonFormat, Name: "body"},
	"$request_time": {Run: CostFormat, Name: "cost"},
	"[$time_local]": {Run: DateFormat, Name: "time"},
}

func DateFormat(s string) interface{} {
	//fmt.Println(s)
	s = s[1 : len(s)-1]
	//s = strings.Replace(s, "2021:", "2021T", 1)
	//s = strings.Replace(s, " +0800", ".000Z", 1)
	//s = strings.Replace(s, "/", "-", 2)
	//s = strings.Trim(s, " ")
	//fmt.Println(s)
	//v, err := time.Parse(DateLayout, s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(s)
	y, _ := strconv.Atoi(s[7:11])
	d, _ := strconv.Atoi(s[0:2])
	h, _ := strconv.Atoi(s[12:14])
	m, _ := strconv.Atoi(s[15:17])
	ss, _ := strconv.Atoi(s[18:20])
	dt := time.Date(y, time.Month(9), d,
		h, m, ss, 0, time.UTC)
	//fmt.Println(dt.Unix())
	return dt
}

func StringFormat(s string) interface{} {
	return s
}

func NumberFormat(s string) interface{} {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func FloatFormat(s string) interface{} {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func CostFormat(s string) interface{} {
	if strings.Index(s, ",") > 0 {
		ls := strings.LastIndex(s, " ")
		s = s[ls+1:]
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func JsonFormat(s string) interface{} {
	var res map[string]interface{}
	ss, err := strconv.Unquote("\"" + s + "\"")
	if err == nil {
		s = ss
	}
	err = json.Unmarshal([]byte(s), &res)

	if err != nil {
		log.Fatal(err)
	}
	return res
}

func ConvertFormat(fm string) []convertor.Convertor {
	fs := strings.Split(fm, " ")
	r := make([]convertor.Convertor, 0, len(fs))
	for _, v := range fs {
		quote := helper.RemoveQuotes(&v)
		cvt := FormatMapper[v]
		cvt.IsQuote = quote
		r = append(r, cvt)
	}
	return r
}
