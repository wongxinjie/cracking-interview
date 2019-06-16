/**
* @File: mininaltree.go
* @Author: wongxinjie
* @Date: 2019/6/11
*/
package main

import (
	"fmt"
	"github.com/wongxinjie/cracking-interview/tree/treenode"
)

func buildTree(array []int, begin, end int) *treenode.BinaryTree {
	if end < begin {
		return nil
	}

	mid := (begin + end) / 2

	node := &treenode.BinaryTree{
		Value: array[mid],
		Left: nil,
		Right: nil,
	}
	node.Left = buildTree(array, begin, mid - 1)
	node.Right = buildTree(array, mid +1, end)

	return node
}

func CreateMininalBST(array []int) *treenode.BinaryTree {
	return buildTree(array, 0, len(array) - 1)
}

func main() {
	t := CreateMininalBST([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(t)
}