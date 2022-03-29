/*
-------------------------------------
# @Time    : 2022/3/29 12:00:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 122_买卖股票的最佳时机 II.go
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
	fmt.Println(maxProfit2(prices))
}

func maxProfit2(prices []int) (ans int) {
	// 贪心
	//for i := 1; i < len(prices); i++ {
	//	if prices[i]-prices[i-1] > 0 {
	//		ans += prices[i] - prices[i-1]
	//	}
	//}
	//return

	// 动态规划
	// dp[i][0]表示当天持有股票时所得现金, dp[i][1]表示当天不持有股票时所得现金
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]-prices[i]) // 保持持有或者今天买入
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i]) // 保持不持有或者今天卖出
	}
	return dp[len(prices)-1][1]
}
