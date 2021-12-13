/*
-------------------------------------
# @Time    : 2021/12/13 19:39:17
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
	fmt.Println(maxProfit122(prices))
}

func maxProfit122(prices []int) (ans int) {
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 {
			ans += prices[i+1] - prices[i]
		}
	}
	return
}
