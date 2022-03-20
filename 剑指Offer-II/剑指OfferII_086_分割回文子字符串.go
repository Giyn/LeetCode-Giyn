/*
-------------------------------------
# @Time    : 2022/3/19 23:48:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_086_分割回文子字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "google"
	fmt.Println(partition(s))
}

func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	var path []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, path...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return
}
