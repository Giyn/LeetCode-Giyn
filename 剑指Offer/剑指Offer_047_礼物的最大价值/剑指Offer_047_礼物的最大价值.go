/*
-------------------------------------
# @Time    : 2022/8/25 0:13:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_047_礼物的最大价值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	var grid = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(grid))
}

func maxValue(grid [][]int) int {
	// 二维动规 每一个dp值只会受到上方格子和左侧格子的影响 所以可以使用一维数组完成
	// dp[j]的上方格子为dp[j] 左侧为dp[j-1]
	dp := make([]int, len(grid[0]))
	for _, row := range grid {
		for j, cur := range row {
			if j == 0 {
				dp[j] += cur
			} else {
				dp[j] = max(dp[j-1], dp[j]) + cur
			}
		}
	}
	return dp[len(dp)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
