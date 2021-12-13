/*
-------------------------------------
# @Time    : 2021/12/13 20:01:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 121_买卖股票的最佳时机.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit121(prices))
}

func maxProfit121(prices []int) (ans int) {
	if len(prices) <= 1 {
		return
	}
	minPrice := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > ans {
			ans = prices[i] - minPrice
		}
	}
	return

	// dp
	//n := len(prices)
	//// dp[i]表示前i天的最大利润 因为始终要使利润最大化
	//dp := make([]int, n)
	//minPrice := prices[0]
	//for i := 1; i < n; i++ {
	//	if prices[i] < minPrice {
	//		minPrice = prices[i]
	//	}
	//	if prices[i]-minPrice > dp[i-1] {
	//		dp[i] = prices[i] - minPrice
	//	} else {
	//		dp[i] = dp[i-1]
	//	}
	//}
	//return dp[len(dp)-1]
}
