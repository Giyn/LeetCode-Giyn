/*
-------------------------------------
# @Time    : 2022/1/29 13:12:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 121_买卖股票的最佳时机.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) (ans int) {
	n := len(prices)
	dp := make([][2]int, n)
	// 0是持有,1是不持有
	dp[0][0] -= prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], -prices[i])
		dp[i][1] = Max(dp[i-1][1], prices[i]+dp[i-1][0])
	}
	return dp[n-1][1]
}
