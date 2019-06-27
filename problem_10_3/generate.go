/**
* @File: generate.go
* @Author: wongxinjie
* @Date: 2019/6/25
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_INT = 2147483648

func generateData(path string) {
	writer, err := os.Create(path)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	
	defer writer.Close()
	for i := 0; i < MAX_INT; i++ {
		if i == 10002 || i == 10004 {
			continue
		}
		num := fmt.Sprintf("%d\n", i)
		_, err = writer.WriteString(num)
		if err != nil {
			fmt.Println("%v", err)
		}
		
	}
}


func Generate(path string) int {
	bitMap := make([]byte, 0, MAX_INT / 8)
	for i := 0; i < MAX_INT / 8 ; i++ {
		bitMap = append(bitMap, byte(0))
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return  -1
	}
	defer file.Close()

	buffer := bufio.NewScanner(file)
	for buffer.Scan() {
		line := strings.TrimSpace(buffer.Text())
		n, _ := strconv.Atoi(line)

		bitMap[ n / 8] |= 1 << uint(n % 8)
	}

	bitMapSize := len(bitMap)
	for i := 0; i < bitMapSize; i++ {
		for j :=0; j < 8; j++ {
			if (bitMap[i] & (1 << uint(j))) == 0 {
				return  i * 8 + j
			}
		}
	}

	return 0
}

func main() {
	r := Generate( "number.txt")
	fmt.Println(r)
}
