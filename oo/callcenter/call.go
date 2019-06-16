/**
* @File: call.go
* @Author: wongxinjie
* @Date: 2019/6/16
*/
package callcenter

import (
	"fmt"
	"time"
)

type Call struct {
	rank int
	caller *Caller
	employee *Employee
	callTime int64
}

func CreateCall(c *Caller) *Call {
	return &Call{
		rank: Responder,
		caller: c,
		employee: nil,
		callTime: time.Now().Unix(),
	}
}

func (c *Call) setHandler(e *Employee) {
	c.employee = e
}


func (c *Call) replay(message string) {
	fmt.Println(message)
}

func (c *Call) incrementRank() int{
	if c.rank == Responder {
		c.rank = Manager
	} else if c.rank == Manager {
		c.rank = Director
	}
	return c.rank
}

func (c *Call) disconnect() {
	fmt.Println("Thank for you calling!")
}