/**
* @File: matrix.go
* @Author: wongxinjie
* @Date: 2019/6/8
*/
package array

func rotate(matrix [][]int, n int) {
	for layer := 0; layer < n / 2; layer++ {
		first := layer
		last := n - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first

			top := matrix[first][i]

			matrix[first][i] = matrix[last-offset][first]
			matrix[last - offset][first] = matrix[last][last - offset]
			matrix[last][last - offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
}


func setZero(matrix [][]int) {
	if len(matrix) <= 0 {
		return
	}

	rowSet := make([]bool, len(matrix))
	columnSet := make([]bool, len(matrix[0]))

	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = true
				columnSet[j] = true
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rowSet[i] || columnSet[j] {
				matrix[i][j] = 0
			}
		}
	}
}