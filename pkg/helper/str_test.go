package helper

import (
	"fmt"
	"testing"
)

func TestRemoveQuotes(t *testing.T) {
	has := "\"a\""
	no := "a"
	type args struct {
		v *string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test-has", args: args{&has}, want: true},
		{name: "test-no", args: args{&no}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveQuotes(tt.args.v); got != tt.want {
				t.Errorf("RemoveQuotes() = %v, want %v", got, tt.want)
			}
		})
	}

	if has != no {
		t.Errorf("RemoveQuotes() %v = %v", has, no)
	} else {
		fmt.Println("check pass")
	}
}
