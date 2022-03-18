/*
-------------------------------------
# @Time    : 2022/2/18 21:50:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 714_买卖股票的最佳时机含手续费.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	fmt.Println(maxProfit6(prices, fee))
}

func maxProfit6(prices []int, fee int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])     // 保持买入状态
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee) // 保持卖出状态
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}
