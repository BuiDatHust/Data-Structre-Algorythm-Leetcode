package RotateImage

import "fmt"

func swapVertex(matrix [][]int, l, r int) {
	temp1 := matrix[l][r]
	matrix[l][r] = matrix[l][l]
	temp2 := matrix[r][r]
	matrix[r][r] = temp1
	temp3 := matrix[r][l]
	matrix[r][l] = temp2
	matrix[l][l] = temp3
}
func swap(matrix [][]int, l, r, i int) {
	temp1 := matrix[i][r]
	matrix[i][r] = matrix[l][i]
	temp2 := matrix[r][r-i]
	matrix[r][r-i] = temp1
	temp3 := matrix[r-i][l]
	matrix[r-i][l] = temp2
	matrix[l][i] = temp3
}

func Rotate1(matrix [][]int) {
	l := 0
	r := len(matrix) - 1
	for l < r {
		for i := l; i < r; i++ {
			if i == l {
				swapVertex(matrix, l, r)
			} else {
				swap(matrix, l, r, i)
			}
		}
		l++
		r--
	}
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val, " ")
		}
		fmt.Print("\n")
	}
}

// good algorythm, time: O(n*n) space: O(1)
// reverse slice then swap i,j <-> j,i
func Rotate2(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][j] = matrix[j][j], matrix[i][j]
		}
	}
}
