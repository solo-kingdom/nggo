package main

import (
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/process"
)

func main() {
	f := "$remote_user"
	fc := process.FormatMapper[f]
	fmt.Printf("%T %s %T\n", fc, fc.Run("hello"), fc.Run("hello"))

	f = "$request_body"
	fc = process.FormatMapper[f]
	s := "{\x22channel\x22:\x22home\x22,\x22uid\x22:\x22692323\x22,\x22log_id\x22:\x2237943eb5415bd698\x22,\x22eid\x22:\x22base\x22,\x22limit\x22:\x2210\x22,\x22top\x22:\x22false\x22}\x0A"
	fmt.Println(fc.Run(s))
}
