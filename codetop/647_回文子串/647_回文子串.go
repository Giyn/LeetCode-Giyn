/*
-------------------------------------
# @Time    : 2022/6/27 1:03:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 647_回文子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) (ans int) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || dp[i+1][j-1]) {
				dp[i][j] = true
				ans++
			}
		}
	}
	return
}
