/**
* @File: treenode.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package treenode

import "fmt"

type BinaryTree struct {
	Value int
	Left *BinaryTree
	Right *BinaryTree
}


func (t *BinaryTree) String() string {
	s := ""

	if t == nil {
		return s
	}

	queue := make([]*BinaryTree, 0)
	queue = append(queue, t)

	for len(queue) > 0 {
		count := len(queue)

		for count > 0 {
			node := queue[0]
			queue = queue[1:]
			s += fmt.Sprintf("%v\t", node.Value)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			count -=1
		}
		s += fmt.Sprint("\n")
	}
	return s
}

