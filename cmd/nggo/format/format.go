package main

import (
	"fmt"
	"github.com/solo-kingdom/nggo.git/pkg/process"
)

func main() {
	f := "$remote_user"
	fc := process.FormatMapper[f]
	fmt.Printf("%T %s %T\n", fc, fc("hello"), fc("hello"))
}
