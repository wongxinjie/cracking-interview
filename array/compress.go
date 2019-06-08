/**
* @File: compress.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

import (
	"strconv"
	"strings"
)

func Compress(str string) string {

	compressedSize := countCompression(str)
	if  compressedSize >= len(str) {
		return str
	}

	var buffer strings.Builder

	length := len(str)
	last := str[0]
	count := 1
	for i := 1; i < length; i++ {
		if str[i] == last {
			count += 1
		} else {
			buffer.WriteString(string(last))
			buffer.WriteString(strconv.FormatInt(int64(count), 10 ))
			last = str[i]
			count = 1
		}
	}

	buffer.WriteString(string(last))
	buffer.WriteString(strconv.FormatInt(int64(count), 10))
	return buffer.String()
}

func countCompression(str string) int {
	if len(str) == 0 {
		return 0
	}

	length := len(str)
	last := str[0]
	size := 0
	count := 1
	for i := 1; i < length; i++ {
		if str[i] == last {
			count += 1
		} else {
			last = str[i]
			size += 1 + len(string(count))
			count = 1
		}
	}

	size += 1 + len(string(count))
	return size
}

