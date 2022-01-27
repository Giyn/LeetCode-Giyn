/*
-------------------------------------
# @Time    : 2022/1/27 18:13:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 279_完全平方数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 13
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt64
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}
