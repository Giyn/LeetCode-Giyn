/*
-------------------------------------
# @Time    : 2021/11/7 13:04:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 598_范围求和 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	m := 3
	n := 3
	ops := [][]int{{2, 2}, {3, 3}}
	fmt.Println(maxCount(m, n, ops))
}

func maxCount(m int, n int, ops [][]int) int {
	minR, minC := m, n
	for _, row := range ops {
		minR = min(minR, row[0])
		minC = min(minC, row[1])
	}
	return minR * minC
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
