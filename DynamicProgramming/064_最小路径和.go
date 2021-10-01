/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 064_最小路径和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	var grid [][]int

	row1 := []int{1, 3, 1}
	row2 := []int{1, 5, 1}
	row3 := []int{4, 2, 1}
	grid = append(grid, row1)
	grid = append(grid, row2)
	grid = append(grid, row3)

	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
