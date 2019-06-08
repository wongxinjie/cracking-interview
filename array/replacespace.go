/**
* @File: replacespace.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

import "strings"

const SPACE = rune(32)

func replaceSpaces(str string) string {
	array := make([]string, 0)

	for _, c := range str {
		if c == SPACE {
			array = append(array, "%20")
		} else {
			array = append(array, string(c))
		}
	}
	return strings.Join(array, "")
}