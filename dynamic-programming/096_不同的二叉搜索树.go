/*
-------------------------------------
# @Time    : 2022/1/22 13:01:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 096_不同的二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 10
	fmt.Println(numTrees(n))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
