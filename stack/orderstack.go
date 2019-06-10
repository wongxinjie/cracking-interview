/**
* @File: orderstack.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package main

import (
	"fmt"
	"github.com/wongxinjie/cracking-interview/stack/stack"
)

func SortedStack(s *stack.Stack) *stack.Stack {
	r := stack.New()

	for !s.IsEmpty() {
		v := s.Pop()
		for !r.IsEmpty() && r.Peek() > v {
			s.Push(r.Pop())
		}
		r.Push(v)
	}
	return r
}


func main() {
	s := stack.New()
	s.Push(1)
	s.Push(4)
	s.Push(5)
	s.Push(3)
	s.Push(2)

	r := SortedStack(s)
	fmt.Println(r.Peek())
}