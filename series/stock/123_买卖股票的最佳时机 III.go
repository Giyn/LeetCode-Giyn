/*
-------------------------------------
# @Time    : 2022/3/29 12:48:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 123_买卖股票的最佳时机 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit3(prices))
}

func maxProfit3(prices []int) int {
	// 0:无操作,1:第一次买入,2:第一次卖出,3:第二次买入,4:第二次卖出
	dp := make([][5]int, len(prices))
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = Max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = Max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = Max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}
