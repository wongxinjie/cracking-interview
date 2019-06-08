/**
* @File: uniquechars.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

func isUniqueChars(str string) bool {
	if len(str) > 256 {
		return false
	}

	charSet := make([]bool, 256)
	for _, c := range str {
		if charSet[int(c)] {
			return false
		}
		charSet[c] = true
	}
	return true
}
