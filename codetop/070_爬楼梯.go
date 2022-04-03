/*
-------------------------------------
# @Time    : 2022/1/19 20:20:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 070_爬楼梯.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 30
	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
