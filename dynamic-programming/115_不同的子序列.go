/*
-------------------------------------
# @Time    : 2022/2/24 9:53:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 115_不同的子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "babgbag"
	t := "bag"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) (ans int) {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
