/**
* @File: stack.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package stack

type Stack struct {
	elements []int64
	count int64
}

func New() *Stack {
	return &Stack{
		make([]int64, 0),
		0,
	}
}

func (s *Stack) Push(v int64) {
	s.elements = append(s.elements, v)
	s.count += 1
}


func (s *Stack) Pop() int64 {
	v := s.elements[s.count - 1]
	s.elements = s.elements[:s.count-1]
	s.count -= 1
	return v
}

func (s *Stack) Peek() int64 {
	return s.elements[s.count-1]
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

