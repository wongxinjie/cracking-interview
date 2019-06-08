/**
* @File: deletedup.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package main

import (
	"fmt"
	"github.com/wongxinjie/cracking-interview/linkedlist/linkedlist"
)

func deleteDup(l *linkedlist.LinkedList) {
	table := make(map[int]bool)

	var prev *linkedlist.LinkedListNode
	current := l.Head

	for current != nil {
		if ok, _ := table[current.Val]; ok {
			prev.Next = current.Next
		} else {
			table[current.Val] = true
			prev = current
		}
		current = current.Next
	}
}


func main() {
	nums := []int{1, 2, 2, 3, 5, 4, 3, 2}
	l := linkedlist.New(nums)
	fmt.Println(l)
	deleteDup(l)
	fmt.Println(l)
}

