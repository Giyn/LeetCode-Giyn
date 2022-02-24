/*
-------------------------------------
# @Time    : 2022/2/24 10:40:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 583_两个字符串的删除操作.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	word1 := "leetcode"
	word2 := "etco"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Min(dp[i-1][j-1]+2, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
