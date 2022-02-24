/*
-------------------------------------
# @Time    : 2022/2/24 9:51:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1706_球会落何处.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}
	fmt.Println(findBall(grid))
}

func findBall(grid [][]int) []int {
	n := len(grid[0])
	ans := make([]int, n)
	for j := range ans {
		col := j // 球的初始列
		for _, row := range grid {
			dir := row[col]
			col += dir                                  // 移动球
			if col < 0 || col == n || row[col] != dir { // 到达侧边或 V 形
				col = -1
				break
			}
		}
		ans[j] = col // col >= 0 为成功到达底部
	}
	return ans
}
