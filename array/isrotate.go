/**
* @File: isrotate.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

import (
	"fmt"
	"strings"
)

func isRotate(x, y string) bool {
	target := fmt.Sprintf("%s%s", x , x)
	return strings.Contains(target, y)
}

