/**
* @File: caller.go
* @Author: wongxinjie
* @Date: 2019/6/16
*/
package callcenter

import "fmt"

type Caller struct {
	name string
	userId int64
	number string
}

func CreateCaller(name string, userId int64, number string) *Caller {
	return &Caller{
		name: name,
		userId: userId,
		number: number,
	}
}

func (caller *Caller) String() string {
	s := fmt.Sprintf("Caller{name: %v, userId: %v, number: %v}",
		caller.name, caller.userId, caller.number)
	return s
}