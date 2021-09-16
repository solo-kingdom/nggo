package process

import (
	"encoding/json"
	"log"
	"strconv"
)

type Convertor func(string) interface{}

//goland:noinspection GoUnusedGlobalVariable
var FormatMapper = map[string]Convertor{
	"$remote_user":  StringFormat,
	"$request_body": JsonFormat,
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

func JsonFormat(s string) interface{} {
	var res map[string]interface{}
	err := json.Unmarshal([]byte(s), &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
