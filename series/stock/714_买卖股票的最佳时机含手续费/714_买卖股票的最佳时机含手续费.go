/*
-------------------------------------
# @Time    : 2022/3/29 16:16:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 714_买卖股票的最佳时机含手续费.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	// 0:持有状态,1:不持有状态
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
