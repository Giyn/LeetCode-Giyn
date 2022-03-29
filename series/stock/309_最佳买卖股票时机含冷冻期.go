/*
-------------------------------------
# @Time    : 2022/3/29 14:09:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 309_最佳买卖股票时机含冷冻期.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit5(prices))
}

func maxProfit5(prices []int) int {
	// 0:买入状态,1:卖出状态,2:今天卖出,3:冷冻期
	dp := make([][4]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return Max(dp[len(prices)-1][1], dp[len(prices)-1][2], dp[len(prices)-1][3])
}
