/**
* @File: addlist.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package main

import (
	"github.com/wongxinjie/cracking-interview/list/linkedlist"
)

func addLists(x, y *linkedlist.LinkedList) *linkedlist.LinkedList {
	var sumHead, current *linkedlist.LinkedListNode

	carry := 0
	count := 0

	xNode := x.Head
	yNode := y.Head

	for xNode != nil && yNode != nil {
		sum := xNode.Val + yNode.Val + carry
		if sum > 10 {
			carry = sum / 10
			sum = sum % 10
		} else {
			carry = 0
		}
		node := &linkedlist.LinkedListNode{sum, nil, nil}
		if sumHead == nil {
			sumHead = node
			current = node
		} else {
			current.Next = node
			current = node
		}

		xNode = xNode.Next
		yNode = yNode.Next

		count += 1
	}

	if xNode != nil {
		sum := xNode.Val + carry
		if sum > 10 {
			carry /= 10
			sum %= sum
		} else {
			carry = 0
		}
		node := &linkedlist.LinkedListNode{sum, nil, nil}
		current.Next = node
		current = node

		xNode = xNode.Next
		count += 1
	}

	if yNode != nil {
		sum := yNode.Val + carry
		if sum > 10 {
			carry /= 10
			sum %= sum
		} else {
			carry = 0
		}
		node := &linkedlist.LinkedListNode{sum, nil, nil}
		current.Next = node
		current = node

		yNode = yNode.Next
		count += 1
	}

	if carry != 0 {
		current.Next = &linkedlist.LinkedListNode{carry, nil, nil}
		count += 1
	}

	return &linkedlist.LinkedList{sumHead, count}
}
