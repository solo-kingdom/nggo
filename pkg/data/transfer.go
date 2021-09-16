package data

import (
	"github.com/solo-kingdom/nggo.git/pkg/convertor"
	"github.com/solo-kingdom/nggo.git/pkg/helper"
	"log"
	"strings"
)

func Transfer(line *string, ffm []convertor.Convertor) (map[string]interface{}, error) {
	r := map[string]interface{}{}

	vs := SplitLine(line)
	//fmt.Println(*line)
	//fmt.Printf("ckpt line; %v", strings.Join(vs, ", "))
	for i, v := range vs {
		if len(ffm) <= i {
			log.Fatalf("[ERROR] unexpect log: i=%v, v=%v, vs=%v, line=%v",
				i, v, strings.Join(vs, "|"), *line)
		}
		cvt := ffm[i]
		quote := false
		if cvt.IsQuote {
			quote = helper.RemoveQuotes(&v)
		}
		if cvt.Run == nil || quote != cvt.IsQuote {
			continue
		}

		t := cvt.Run(v)
		switch t.(type) {
		case map[string]interface{}:
			for tk, tv := range t.(map[string]interface{}) {
				r[tk] = tv
			}
		default:
			r[cvt.Name] = t
		}
	}
	return r, nil
}

func SplitLine(s *string) []string {
	var r = make([]string, 0, 15)
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == ' ' {
			continue
		}

		var n uint8 = ' '
		if (*s)[i] == '"' {
			n = '"'
		} else if (*s)[i] == '[' {
			n = ']'
		} else if (*s)[i] == '{' {
			n = '}'
		}
		// TODO
		j := helper.IndexChar(s, i+1, n)
		for j > i && (*s)[j-1] == ',' { // template fix for retry request
			j = helper.IndexChar(s, j+1, n)
		}
		if j < 0 {
			j = len(*s) - 1
		}
		r = append(r, (*s)[i:j+1])
		i = j
	}
	return r
}
