/*
-------------------------------------
# @Time    : 2022/1/30 10:56:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 122_买卖股票的最佳时机 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) (ans int) {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) // 持有(保持持有or今天买入)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) // 不持有(保持不持有or今天卖出)
	}
	return dp[n-1][1]
}
