package process

import (
	"encoding/json"
	"github.com/solo-kingdom/nggo.git/pkg/convertor"
	"github.com/solo-kingdom/nggo.git/pkg/helper"
	"log"
	"strconv"
	"strings"
)

//goland:noinspection GoUnusedGlobalVariable
var FormatMapper = map[string]convertor.Convertor{
	"$remote_user":  {Run: StringFormat, Name: "user"},
	"$request_body": {Run: JsonFormat, Name: "body"},
	"$request_time": {Run: CostFormat, Name: "cost"},
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
