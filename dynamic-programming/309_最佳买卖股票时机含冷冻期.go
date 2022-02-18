/*
-------------------------------------
# @Time    : 2022/2/18 12:42:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 309_最佳买卖股票时机含冷冻期.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfitCool(prices))
}

func maxProfitCool(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i])) // 达到买入股票状态
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])                                      // 达到保持卖出股票状态
		dp[i][2] = dp[i-1][0] + prices[i]                                           // 达到今天卖出了股票状态
		dp[i][3] = dp[i-1][2]                                                       // 达到冷冻期状态
	}
	return max(dp[len(prices)-1][3], max(dp[len(prices)-1][1], dp[len(prices)-1][2]))
}
