package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	s :=
		"{\x22section_id\x22:\x22home_rec\x22,\x22distinct_id\x22:\x221067590\x22,\x22log_id\x22:\x227f51eb67cdd99242\x22,\x22exp_id\x22:\x22sensors_rec\x22,\x22limit\x22:\x2210\x22,\x22need_top_item\x22:\x22false\x22}\x0A"
	//r, err := strconv.Unquote("\"" + s + "\"")

	fmt.Print(s)
	rs := fmt.Sprintf("%s", s)
	fmt.Println(rs)

	//r, err := strconv.Unquote(s)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(r)

	//rs := fmt.Sprintf("%s", s)
	//fmt.Println(rs)
	var m map[string]interface{}
	err := json.Unmarshal([]byte(rs), &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}
