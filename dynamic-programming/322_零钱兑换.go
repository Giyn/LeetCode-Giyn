/*
-------------------------------------
# @Time    : 2022/1/27 16:57:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 322_零钱兑换.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = math.MaxInt64
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt64 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
