/*
-------------------------------------
# @Time    : 2022/1/27 14:48:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 518_零钱兑换 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
