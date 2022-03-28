/*
-------------------------------------
# @Time    : 2022/3/28 10:05:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 048_旋转图像.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
