/*
-------------------------------------
# @Time    : 2022/2/22 13:51:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 392_判断子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	// 双指针
	//i := 0
	//for j := 0; j < len(t) && i < len(s); j++ {
	//	if s[i] == t[j] {
	//		i++
	//	}
	//}
	//return i == len(s)

	// DP
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1] // 相当于t要删除元素,看s[i-1]与t[j-2]的比较结果
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}
