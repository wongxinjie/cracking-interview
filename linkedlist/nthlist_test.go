/**
* @File: nthlist.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package main

import (
	"testing"

	"github.com/wongxinjie/cracking-interview/linkedlist/linkedlist"
)

func Test_ntplist(t *testing.T) {
	type args struct {
		l *linkedlist.LinkedList
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{linkedlist.New([]int{1, 2, 3, 4, 5, 6}), 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ntplist(tt.args.l, tt.args.k); got != tt.want {
				t.Errorf("ntplist() = %v, want %v", got, tt.want)
			}
		})
	}
}
