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

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) (ans int) {
	n := len(prices)
	dp := make([]int, n)
	min := prices[0]
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		dp[i] = max(dp[i-1], prices[i]-min)
	}
	return dp[len(dp)-1]

	//max := func(x, y int) int {
	//	if x > y {
	//		return x
	//	}
	//	return y
	//}
	//n := len(prices)
	//dp := make([][2]int, n)
	//dp[0][0] -= prices[0]
	//dp[0][1] = 0
	//for i := 1; i < n; i++ {
	//	dp[i][0] = max(dp[i-1][0], -prices[i])
	//	dp[i][1] = max(dp[i-1][1], prices[i]+dp[i-1][0])
	//}
	//return dp[n-1][1]
}
