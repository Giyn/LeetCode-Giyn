/*
-------------------------------------
# @Time    : 2022/2/17 0:43:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 688_骑士在棋盘上的概率.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 3
	k := 2
	row := 0
	column := 0
	fmt.Println(knightProbability(n, k, row, column))
}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make([][][]float64, n)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, k+1)
			dp[i][j][0] = 1
		}
	}
	for i := 1; i <= k; i++ {
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				for _, dir := range [][]int{{1, 2}, {2, 1}, {-2, 1}, {1, -2}, {2, -1}, {-1, 2}, {-1, -2}, {-2, -1}} {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < n {
						dp[r][c][i] += dp[nr][nc][i-1] / 8
					}
				}
			}
		}
	}
	return dp[row][column][k]
}
