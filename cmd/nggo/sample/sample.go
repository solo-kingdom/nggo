package main

import (
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/helper"
)

func main() {
	s := "\"a\""
	fmt.Println(helper.RemoveQuotes(&s), s)

	m := map[string]string{}
	fmt.Printf("%v, %T\n", m["A"], m["A"])
}
