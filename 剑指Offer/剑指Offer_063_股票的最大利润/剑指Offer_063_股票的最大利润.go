/*
-------------------------------------
# @Time    : 2022/3/29 11:57:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_063_股票的最大利润.go
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
