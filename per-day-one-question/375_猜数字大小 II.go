/*
-------------------------------------
# @Time    : 2021/11/12 0:01:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 375_猜数字大小 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	fmt.Println(getMoneyAmount(n))
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt64
			for k := i; k < j; k++ {
				cost := k + max375(dp[i][k-1], dp[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			dp[i][j] = minCost
		}
	}
	return dp[1][n]
}

func max375(x, y int) int {
	if x > y {
		return x
	}
	return y
}
