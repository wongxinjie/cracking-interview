/**
* @File: machine.go
* @Author: wongxinjie
* @Date: 2019/6/25
*/
package problem_10_2

type Machine struct {
	persons map[int64]*Person
	machineId int64
}

func CreateMachine(id int64) *Machine {
	return &Machine{
		persons: make(map[int64]*Person, 0),
		machineId: id,
	}
}

func (m *Machine) GetPersonWithId(id int64) *Person {
	return m.persons[id]
}