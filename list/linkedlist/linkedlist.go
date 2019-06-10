/**
* @File: list.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package linkedlist

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	Val int
	Next *LinkedListNode
	Prev *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Count int
}

func New(nums []int) *LinkedList {
	if len(nums) == 0 {
		return nil
	}
	
	count := len(nums)
	head := &LinkedListNode{nums[0], nil, nil}
	current := head
	for i := 1; i < count; i++ {
		node := &LinkedListNode{nums[i], nil, current}
		current.Next = node
		current = node
	}
	
	return &LinkedList{head, count}
}

func (l LinkedList) String() string {
	nodes := make([]string, 0)

	for n := l.Head; n != nil; n = n.Next {
		nodes = append(nodes, strconv.FormatInt(int64(n.Val), 10))
	}
	s := strings.Join(nodes, "->")
	return s
}