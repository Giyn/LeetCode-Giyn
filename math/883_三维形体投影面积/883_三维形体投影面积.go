/*
-------------------------------------
# @Time    : 2022/4/26 21:21:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 883_三维形体投影面积.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println(projectionArea(grid))
}

func projectionArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	xy, xz, yz := 0, 0, 0
	for i := 0; i < m; i++ {
		rowMax, colMax := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				xy++
			}
			rowMax = max(rowMax, grid[i][j])
			colMax = max(colMax, grid[j][i]) // m == n
		}
		xz += rowMax
		yz += colMax
	}
	return xy + xz + yz
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
