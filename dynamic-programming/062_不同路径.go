/*
-------------------------------------
# @Time    : 2022/1/20 20:43:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 062_不同路径.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	m, n := 3, 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
