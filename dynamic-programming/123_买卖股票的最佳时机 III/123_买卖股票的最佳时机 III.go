/*
-------------------------------------
# @Time    : 2022/1/30 12:04:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 123_买卖股票的最佳时机 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	//dp := [5]int{} // 5种状态(0.无操作 1.第一次买入 2.第一次卖出 3.第二次买入 4.第二次卖出)
	//dp[1], dp[3] = -prices[0], -prices[0]
	//for i := 1; i < n; i++ {
	//	dp[1] = max(dp[0]-prices[i], dp[1])
	//	dp[2] = max(dp[1]+prices[i], dp[2])
	//	dp[3] = max(dp[2]-prices[i], dp[3])
	//	dp[4] = max(dp[3]+prices[i], dp[4])
	//}
	//return dp[4]

	dp := make([][5]int, n) // 5种状态(0.无操作 1.第一次买入 2.第一次卖出 3.第二次买入 4.第二次卖出)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1]) // 今天买入or保持昨天第一次买入的状态
		dp[i][2] = max(dp[i-1][1]+prices[i], dp[i-1][2]) // 今天卖出or保持昨天第一次卖出的状态
		dp[i][3] = max(dp[i-1][2]-prices[i], dp[i-1][3])
		dp[i][4] = max(dp[i-1][3]+prices[i], dp[i-1][4])
	}
	return dp[n-1][4]
}
