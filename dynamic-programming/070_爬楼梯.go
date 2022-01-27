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
	//dp := []int{0, 1, 2}
	//if n <= 2 {
	//	return n
	//}
	//for i := 2; i < n; i++ {
	//	tmp := dp[1] + dp[2]
	//	dp[1] = dp[2]
	//	dp[2] = tmp
	//}
	//return dp[2]

	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= 2; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}
		}
	}
	return dp[n]
}
