/**
* @File: ispalindrome.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package main

import (
	"container/list"
	"github.com/wongxinjie/cracking-interview/list/linkedlist"
)

func isPalidrome(l *linkedlist.LinkedList) bool {
	fast := l.Head
	slow := l.Head

	stack := list.New()

	for fast != nil && fast.Next != nil {
		stack.PushBack(slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		top := stack.Back()
		if top.Value != slow.Val {
			return false
		}
		stack.Remove(top)
		slow = slow.Next
	}
	return true
}