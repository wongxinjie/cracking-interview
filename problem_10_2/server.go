/**
* @File: server.go
* @Author: wongxinjie
* @Date: 2019/6/25
*/
package problem_10_2

type Server struct {
	machines map[int64] *Machine
	personMachineMap map[int64]int64
}