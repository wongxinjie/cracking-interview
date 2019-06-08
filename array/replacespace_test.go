/**
* @File: replacespace.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package array

import "testing"

func Test_replaceSpaces(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"replace", args{"a b  c"}, "a%20b%20%20c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceSpaces(tt.args.str); got != tt.want {
				t.Errorf("replaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
