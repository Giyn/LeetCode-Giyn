/*
-------------------------------------
# @Time    : 2022/1/20 17:51:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 746_使用最小花费爬楼梯.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	var min func(int, int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	n := len(cost)
	//dp := make([]int, n)
	//dp[0] = cost[0]
	//dp[1] = cost[1]
	//for i := 2; i < n; i++ {
	//	dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	//}
	//return min(dp[n-1], dp[n-2])

	dp0, dp1 := cost[0], cost[1]
	for i := 2; i < n; i++ {
		dpi := min(dp0, dp1) + cost[i]
		dp0 = dp1
		dp1 = dpi
	}
	return min(dp0, dp1)
}
