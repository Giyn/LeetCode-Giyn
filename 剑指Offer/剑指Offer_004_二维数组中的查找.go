/*
-------------------------------------
# @Time    : 2022/3/22 9:46:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_004_二维数组中的查找.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := -10
	fmt.Println(findNumberIn2DArray(matrix, target))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}
