/*
-------------------------------------
# @Time    : 2022/3/20 15:57:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_091_粉刷房子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	fmt.Println(minCost(costs))
}

func minCost(costs [][]int) int {
	dp := make([]int, 3)
	dp[0] = costs[0][0]
	dp[1] = costs[0][1]
	dp[2] = costs[0][2]
	for i := 1; i < len(costs); i++ {
		tmp := make([]int, 3)
		tmp[0] = min(dp[1], dp[2]) + costs[i][0]
		tmp[1] = min(dp[0], dp[2]) + costs[i][1]
		tmp[2] = min(dp[0], dp[1]) + costs[i][2]
		dp = tmp
	}
	return min(min(dp[0], dp[1]), dp[2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
