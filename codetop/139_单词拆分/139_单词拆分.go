/*
-------------------------------------
# @Time    : 2022/6/27 14:13:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 139_单词拆分.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	mp := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		mp[word] = true
	}
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mp[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
