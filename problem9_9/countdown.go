/**
* @File: countdown.go
* @Author: wongxinjie
* @Date: 2019/6/23
*/
package main

import "fmt"

func countDownDP(n int, cache map[int]int) int{
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if cache[n] > -1 {
		return cache[n]
	} else {
		cache[n] = countDownDP(n -1, cache) + countDownDP(n - 2, cache) + countDownDP(n - 3, cache)
		return cache[n]
	}
}


func main() {
	n := 5
	cache := make(map[int]int, 0)
	for i := 0; i < n+1; i ++ {
		cache[i] =  -1
	}
	fmt.Println(countDownDP(n, cache))
}