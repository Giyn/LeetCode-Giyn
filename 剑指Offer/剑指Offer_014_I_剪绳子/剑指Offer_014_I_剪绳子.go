/*
-------------------------------------
# @Time    : 2022/3/31 17:09:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_014_I_剪绳子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	n := 10
	fmt.Println(cuttingRope(n))
}

func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(max(dp[i], (i-j)*j), dp[i-j]*j)
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
