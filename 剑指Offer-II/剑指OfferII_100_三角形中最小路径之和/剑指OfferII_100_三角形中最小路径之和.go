/*
-------------------------------------
# @Time    : 2022/3/21 8:25:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_100_三角形中最小路径之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) (ans int) {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]

	//n := len(triangle)
	//dp := make([][]int, n+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	//for i := n - 1; i >= 0; i-- {
	//	for j := 0; j <= i; j++ {
	//		dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
	//	}
	//}
	//return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
