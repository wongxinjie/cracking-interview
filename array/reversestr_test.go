/**
* @File: reversestr.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package array

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"even", args{"rule"}, "elur"},
		{"odd", args{"thing"}, "gniht"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.str); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
