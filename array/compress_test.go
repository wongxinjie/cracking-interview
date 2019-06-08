/**
* @File: compress.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package array

import "testing"

func TestCompress(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"compress", args{"aaaabbccc"}, "a4b2c3"},
		{"no compressed", args{"abc"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compress(tt.args.str); got != tt.want {
				t.Errorf("Compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countCompression(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{"abc"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompression(tt.args.str); got != tt.want {
				t.Errorf("countCompression() = %v, want %v", got, tt.want)
			}
		})
	}
}
