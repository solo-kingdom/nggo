package data

import (
	"fmt"
	"strings"
)

func Transfer(line string, fm string) map[string]string {
	fs := strings.Split(fm, " ")
	fmt.Println(fs, len(fs))
	return nil
}