/*
-------------------------------------
# @Time    : 2021/10/25 0:01:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 240_搜索二维矩阵 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	//matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	matrix := [][]int{{-5}}
	//matrix := [][]int{{5},{6}}
	//matrix := [][]int{{1, 4}, {2, 5}}
	target := -5
	fmt.Println(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}
	}
	return false
}
