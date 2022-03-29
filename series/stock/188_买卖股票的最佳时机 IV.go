/*
-------------------------------------
# @Time    : 2022/3/29 13:47:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 188_买卖股票的最佳时机 IV.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit4(k, prices))
}

func maxProfit4(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// i+1为第i次买入,i+2为第i次卖出
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for j := 1; j < 2*k; j += 2 {
		dp[0][j] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = Max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = Max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}
