/**
* @File: call.go
* @Author: wongxinjie
* @Date: 2019/6/17
*/
package main

import "github.com/wongxinjie/cracking-interview/oo/callcenter"

func main() {
	d := callcenter.InitDispatcher(3, 1, 1, 1)
	d.ReceiveComingCall(nil)
}
