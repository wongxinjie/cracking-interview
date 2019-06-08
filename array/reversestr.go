/**
* @File: reversestr.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

func reverse(str string) string {
	chars := []rune(str)

	end := len(chars) - 1
	begin := 0
	for begin < end {
		chars[begin], chars[end] = chars[end], chars[begin]
		begin += 1
		end -= 1
	}
	return string(chars)
}
