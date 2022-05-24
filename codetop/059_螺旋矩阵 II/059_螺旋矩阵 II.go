/*
-------------------------------------
# @Time    : 2022/5/23 9:35:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 059_螺旋矩阵 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	count := 1
	sum := n * n
	for count <= sum {
		for i := left; i <= right; i++ {
			if count <= sum {
				matrix[top][i] = count
				count++
			}
		}
		top++
		for i := top; i <= bottom; i++ {
			if count <= sum {
				matrix[i][right] = count
				count++
			}
		}
		right--
		for i := right; i >= left; i-- {
			if count <= sum {
				matrix[bottom][i] = count
				count++
			}
		}
		bottom--
		for i := bottom; i >= top; i-- {
			if count <= sum {
				matrix[i][left] = count
				count++
			}
		}
		left++
	}
	return matrix
}
