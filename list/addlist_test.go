/**
* @File: addlist.go
* @Author: wongxinjie
* @Date: 2019/6/9
 */
package main

import (
	"reflect"
	"testing"

	"github.com/wongxinjie/cracking-interview/list/linkedlist"
)

func Test_addLists(t *testing.T) {
	type args struct {
		x *linkedlist.LinkedList
		y *linkedlist.LinkedList
	}
	tests := []struct {
		name string
		args args
		want *linkedlist.LinkedList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addLists(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
