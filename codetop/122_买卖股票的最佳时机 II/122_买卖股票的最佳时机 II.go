/*
-------------------------------------
# @Time    : 2022/4/26 17:40:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 122_买卖股票的最佳时机 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	n := len(prices)
	// 0: 持有, 1: 不持有
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
