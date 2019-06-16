/**
* @File: main.go
* @Author: wongxinjie
* @Date: 2019/6/10
*/
package main

import (
	"fmt"
	"github.com/wongxinjie/cracking-interview/graph/graph"
)

func main() {
	m := map[int64][]int64{
		12: {1, 2},
		1: {2},
		2: {12},
	}
	fmt.Println(graph.CreateGraphFromMap(m))
}
