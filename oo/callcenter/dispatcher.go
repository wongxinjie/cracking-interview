/**
* @File: dispatcher.go
* @Author: wongxinjie
* @Date: 2019/6/16
*/
package callcenter

import "fmt"

var DISPATCHER *Dispatcher

type Dispatcher struct {
	// 雇员层级 接线员、主管和经理
	level int
	respondentCount int
	managerCount int
	directorCount int
	/*
	employees[0] - responders
	employees[1] - managers
	employees[2] -
	 */
	employeesLevel [][]*Employee
	callQueue [][]*Call

}

func InitDispatcher(level, respondentCount, managerCount, directorCount int) *Dispatcher{
	employees := make([][]*Employee, 0)

	responders := make([]*Employee, 0)
	for i := 0; i < respondentCount; i++ {
		r := CreateEmployeee(Responder)
		responders = append(responders, r)
	}
	employees = append(employees, responders)

	managers := make([]*Employee, 0)
	for i := 0; i < managerCount; i++ {
		m := CreateEmployeee(Manager)
		managers = append(managers, m)
	}
	employees = append(employees, managers)

	directors := make([]*Employee, 0)
	for i := 0; i < directorCount; i++ {
		d := CreateEmployeee(Director)
		directors = append(directors, d)
	}
	employees = append(employees, directors)

	calls := make([][]*Call, 0)
	for i := 0; i < level; i ++ {
		c := make([]*Call, 0)
		calls = append(calls, c)
	}

	dispatcher := &Dispatcher{
		level: level,
		respondentCount: respondentCount,
		managerCount: managerCount,
		directorCount: directorCount,

		employeesLevel: employees,
		callQueue: calls,
	}
	DISPATCHER = dispatcher
	return dispatcher
}


func (d *Dispatcher) getHandlerForCall(c *Call) *Employee {
	for r := c.rank; r < d.level; r++ {
		employees := d.employeesLevel[r]
		for _, e := range employees {
			if e.isFree() {
				return e
			}
		}
	}
	return nil
}


func (d *Dispatcher) ReceiveComingCall(caller *Caller) {
	call := CreateCall(caller)
	fmt.Printf("Receive-Call: name->%v, number-%v, time-%v\n", caller.name, caller.number, call.callTime)
	d.dispatcherCall(call)
}

func (d *Dispatcher) dispatcherCall(call *Call) {
	e := d.getHandlerForCall(call)
	if e != nil {
		e.receivedCall(call)
		call.setHandler(e)
		//e.callCompleted()
	} else {
		call.replay("Plz wait for free employee to reply")
		queue := d.callQueue[call.rank]
		queue = append(queue, call)
	}
}

func (d *Dispatcher) assignCall(e *Employee) bool {
	for r := e.rank; r >= 0; r-- {
		calls := d.callQueue[r]

		if len(calls) > 0 {
			call := calls[0]
			d.callQueue[r] = calls[1:]
			if call != nil {
				e.receivedCall(call)
				return true
			}
		}
	}
	return false
}

