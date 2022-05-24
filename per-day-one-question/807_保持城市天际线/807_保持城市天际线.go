/*
-------------------------------------
# @Time    : 2021/12/13 8:53:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 807_保持城市天际线.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) (ans int) {
	var rowMax = make([]int, len(grid))
	var colMax = make([]int, len(grid[0]))
	for i, row := range grid {
		for j, h := range row {
			rowMax[i] = max(rowMax[i], h)
			colMax[j] = max(colMax[j], h)
		}
	}
	for i, row := range grid {
		for j, h := range row {
			ans += min(rowMax[i], colMax[j]) - h
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
