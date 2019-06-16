/**
* @File: isbalance.go
* @Author: wongxinjie
* @Date: 2019/6/9
*/
package main

import (
	"fmt"
	"math"

	"github.com/wongxinjie/cracking-interview/tree/treenode"
)

func IsBalance(t *treenode.BinaryTree) bool {
	if checkHeight(t) == -1 {
		return false
	}

	return true
}


func checkHeight(root *treenode.BinaryTree) int {
	if root == nil {
		return 0
	}

	leftHeight := checkHeight(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := checkHeight(root.Right)
	if rightHeight == -1 {
		return -1
	}

	diffHeight := int(math.Abs(float64(leftHeight) - float64(rightHeight)))
	if diffHeight > 1 {
		return -1
	}

	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}


func main() {
	root := treenode.BinaryTree{1, nil, nil }
	root.Left = &treenode.BinaryTree{2, nil, nil}
	root.Right = &treenode.BinaryTree{3, nil, nil }

	root.Left.Left = &treenode.BinaryTree{1, nil, nil}
	root.Left.Left.Left = &treenode.BinaryTree{4, nil, nil}
	fmt.Println(IsBalance(&root))
}
