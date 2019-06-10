/**
* @File: ispalindrome.go
* @Author: wongxinjie
* @Date: 2019/6/9
 */
package main

import (
	"testing"

	"github.com/wongxinjie/cracking-interview/list/linkedlist"
)

func Test_isPalidrome(t *testing.T) {
	type args struct {
		l *linkedlist.LinkedList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{linkedlist.New([]int{1, 2, 3, 2, 1})}, true},
		{"no", args{linkedlist.New([]int{1, 2, 3, 4})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalidrome(tt.args.l); got != tt.want {
				t.Errorf("isPalidrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
