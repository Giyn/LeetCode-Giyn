/*
-------------------------------------
# @Time    : 2022/4/8 21:37:52
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
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) (ans int) {
	min := math.MaxInt64
	for _, price := range prices {
		if price < min {
			min = price
		}
		ans = max(ans, price-min)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
