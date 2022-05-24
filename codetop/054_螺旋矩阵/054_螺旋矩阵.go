/*
-------------------------------------
# @Time    : 2022/4/26 12:31:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 054_螺旋矩阵.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	total := m * n
	cnt := 0
	left, right := 0, n-1
	top, bottom := 0, m-1
	for cnt < total {
		for i := left; i <= right; i++ {
			if cnt < total {
				ans = append(ans, matrix[top][i])
				cnt++
			} else {
				break
			}
		}
		top++
		for i := top; i <= bottom; i++ {
			if cnt < total {
				ans = append(ans, matrix[i][right])
				cnt++
			} else {
				break
			}
		}
		right--
		for i := right; i >= left; i-- {
			if cnt < total {
				ans = append(ans, matrix[bottom][i])
				cnt++
			} else {
				break
			}
		}
		bottom--
		for i := bottom; i >= top; i-- {
			if cnt < total {
				ans = append(ans, matrix[i][left])
				cnt++
			} else {
				break
			}
		}
		left++
	}
	return
}
