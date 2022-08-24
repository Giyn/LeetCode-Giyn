/*
-------------------------------------
# @Time    : 2022/8/25 0:24:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_060_n个骰子的点数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 2
	fmt.Println(dicesProbability(n))
}

func dicesProbability(n int) []float64 {
	// i代表当前有几个骰子 j代表当前拼成的数字总和是多少 如果现在的总和是j 那么最后一个骰子可能分别摇到 1 2 3 4 5 6
	// dp[i][j] += for (1~6) dp[i-1][j-k] (k代表了1到6)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i + 1; j <= 6*i; j++ {
			for k := 1; k <= 6; k++ {
				if j-k >= i-1 {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}

	result := make([]float64, 6*n)
	for i := n; i <= 6*n; i++ {
		result[i-1] = float64(dp[n][i]) / math.Pow(6, float64(n))
	}
	return result[n-1:]
}
