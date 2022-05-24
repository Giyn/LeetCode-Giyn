/*
-------------------------------------
# @Time    : 2021/10/22 2:29:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_029_顺时针打印矩阵.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
	}
	matrix[0] = []int{1, 2, 3, 4}
	matrix[1] = []int{5, 6, 7, 8}
	matrix[2] = []int{9, 10, 11, 12}

	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) (ans []int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	amount := m * n
	cnt := 0
	top, bottom := 0, m-1
	left, right := 0, n-1

	for cnt < amount {
		for i := left; i <= right; i++ {
			if cnt < amount {
				ans = append(ans, matrix[top][i])
				cnt++
			} else {
				return
			}
		}
		top++
		for i := top; i <= bottom; i++ {
			if cnt < amount {
				ans = append(ans, matrix[i][right])
				cnt++
			} else {
				return
			}
		}
		right--
		for i := right; i >= left; i-- {
			if cnt < amount {
				ans = append(ans, matrix[bottom][i])
				cnt++
			} else {
				return
			}
		}
		bottom--
		for i := bottom; i >= top; i-- {
			if cnt < amount {
				ans = append(ans, matrix[i][left])
				cnt++
			} else {
				return
			}
		}
		left++
	}
	return ans
}
