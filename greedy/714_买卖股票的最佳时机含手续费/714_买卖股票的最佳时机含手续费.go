/*
-------------------------------------
# @Time    : 2021/12/25 20:19:18
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
	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) (ans int) {
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		//if prices[i] >= minPrice && prices[i] <= minPrice+fee {
		//	continue
		//}
		if prices[i] > minPrice+fee {
			ans += prices[i] - minPrice - fee
			minPrice = prices[i] - fee // 关键一步 不是真正卖出 相当于持有股票
		}
	}
	return
}
