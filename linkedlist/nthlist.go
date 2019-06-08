/**
* @File: nthlist.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package main

import "github.com/wongxinjie/cracking-interview/linkedlist/linkedlist"


func ntplist(l *linkedlist.LinkedList, k int) int{

	p1 := l.Head
	p2 := p1

	for i := 0; i < k; i++ {
		if p2 == nil {
			return -1
		}
		p2 = p2.Next
	}

	if p2 == nil {
		return -1
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1.Val
}