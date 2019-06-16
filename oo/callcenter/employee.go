/**
* @File: employee.go
* @Author: wongxinjie
* @Date: 2019/6/16
*/
package callcenter

import "fmt"

type Employee struct {
	currentCall *Call
	rank int
}

func CreateEmployeee(rank int) *Employee {
	return &Employee{
		currentCall: nil,
		rank: rank,
	}
}

func (e *Employee) receivedCall(c *Call) {
	e.currentCall = c
	fmt.Printf("Reply-For-Call: %v\n", c.caller.number)
}

func (e *Employee) callCompleted() {
	if e.currentCall != nil {
		e.currentCall.disconnect()
		e.currentCall = nil
	}
	DISPATCHER.assignCall(e)
}



func (e *Employee) isFree() bool{
	return e.currentCall == nil
}
