/*
-------------------------------------
# @Time    : 2022/2/15 1:23:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1380_矩阵中的幸运数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := [][]int{
		{36376, 85652, 21002, 4510},
		{68246, 64237, 42962, 9974},
		{32768, 97721, 47338, 5841},
		{55103, 18179, 79062, 46542},
	}
	fmt.Println(luckyNumbers(matrix))
}

func luckyNumbers(matrix [][]int) (ans []int) {
	minRow := func(nums []int) (idx int) {
		min := nums[0]
		for i, num := range nums {
			if num < min {
				min = num
				idx = i
			}
		}
		return
	}
	maxCol := func(idx int) (ans int) {
		ans = matrix[0][idx]
		for i := 1; i < len(matrix); i++ {
			if matrix[i][idx] > ans {
				ans = matrix[i][idx]
			}
		}
		return
	}
	for i := 0; i < len(matrix); i++ {
		minRowIndex := minRow(matrix[i])
		if matrix[i][minRowIndex] == maxCol(minRowIndex) {
			ans = append(ans, matrix[i][minRowIndex])
		}
	}
	return
}
