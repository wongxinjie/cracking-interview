/**
* @File: uniquechars.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package array

import "testing"

func Test_isUniqueChars(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{"abcde"}, true},
		{"no", args{"abcdeab"}, false},
		{"no,empty", args{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUniqueChars(tt.args.str); got != tt.want {
				t.Errorf("isUniqueChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
