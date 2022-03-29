/*
-------------------------------------
# @Time    : 2022/3/29 11:55:45
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
	fmt.Println(maxProfit1(prices))
}

func maxProfit1(prices []int) (ans int) {
	minPrice := math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > ans {
			ans = prices[i] - minPrice
		}
	}
	return
}
