/*
-------------------------------------
# @Time    : 2022/1/22 11:50:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 343_整数拆分.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	n := 10
	fmt.Println(integerBreak(n))
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = Max(dp[i], (i-j)*j, dp[i-j]*j)
		}
	}
	return dp[n]
}
