/**
* @File: isbst.go
* @Author: wongxinjie
* @Date: 2019/6/16
*/
package main

import (
	"fmt"
	"github.com/wongxinjie/cracking-interview/common"
	"github.com/wongxinjie/cracking-interview/tree/treenode"
)

func isBST(t *treenode.BinaryTree, min, max int) bool {
	if t == nil {
		return true
	}

	if t.Value < min || t.Value > max {
		return false
	}

	if !isBST(t.Left, min, t.Value) || !isBST(t.Right, t.Value, max) {
		return false
	}
	return true
}

func CheckBST(t *treenode.BinaryTree) bool {
	return isBST(t, common.MIN_INT, common.MAX_INT)
}

func main() {
	t := &treenode.BinaryTree{3, nil, nil}
	t.Left = &treenode.BinaryTree{2, nil, nil}
	t.Right = &treenode.BinaryTree{4, nil, nil }
	fmt.Println(CheckBST(t))
}