/*
-------------------------------------
# @Time    : 2022/4/17 20:18:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 005_最长回文子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) (ans string) {
	begin := 0
	max := math.MinInt64
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1] == true) {
				dp[i][j] = true
				if j-i+1 > max {
					max = j - i + 1
					begin = i
				}
			}
		}
	}
	return s[begin : begin+max]
}
