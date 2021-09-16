package helper

import "strings"

func RemoveQuotes(v *string) bool {
	if strings.HasPrefix(*v, "\"") && strings.HasSuffix(*v, "\"") {
		*v = (*v)[1 : len(*v)-1]
		return true
	}
	return false
}

func IndexChar(v *string, start int, c uint8) int {
	for start < len(*v) {
		if (*v)[start] == c {
			return start
		} else {
			start++
		}
	}
	return -1
}
