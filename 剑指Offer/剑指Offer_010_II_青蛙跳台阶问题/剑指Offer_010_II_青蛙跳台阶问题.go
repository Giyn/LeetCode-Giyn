/*
-------------------------------------
# @Time    : 2022/3/24 10:07:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_010_II_青蛙跳台阶问题.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 7
	fmt.Println(numWays(n))
}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-2] + dp[i-1]) % 1000000007
	}
	return dp[n]
}
