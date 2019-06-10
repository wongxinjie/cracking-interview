/**
* @File: deletedup.go
* @Author: wongxinjie
* @Date: 2019/6/8
 */
package main

import (
	"testing"

	"github.com/wongxinjie/cracking-interview/list/linkedlist"
)

func Test_deleteDup(t *testing.T) {
	type args struct {
		l *linkedlist.LinkedList
	}
	tests := []struct {
		name string
		args args
		result *linkedlist.LinkedList
	}{
		{"ok",
		args{linkedlist.New([]int{1, 2, 3, 4, 2, 3, 4})},
		 linkedlist.New([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteDup(tt.args.l)
			x := tt.args.l.Head
			y := tt.result.Head
			for x != nil && y != nil {
				if x.Val != y.Val {
					t.Error("remove fails")
				}
				x = x.Next
				y = y.Next
			}
		})
	}
}
