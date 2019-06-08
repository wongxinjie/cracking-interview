/**
* @File: isrotate.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package array

import "testing"

func Test_isRotate(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{"watterbottle", "erbottlewat"}, true},
		{"no", args{"xyz", "zz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRotate(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("isRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
