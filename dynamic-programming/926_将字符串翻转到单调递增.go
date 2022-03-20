/*
-------------------------------------
# @Time    : 2022/3/20 17:45:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 926_将字符串翻转到单调递增.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	s := "00011000"
	fmt.Println(minFlipsMonoIncr(s))
}

func minFlipsMonoIncr(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// dp[i][0]表示前i个元素递增且第i个元素为0的最小翻转次数
	// dp[i][1]表示前i个元素递增且第i个元素为1的最小翻转次数
	if s[0] == '1' {
		dp[0][0] = 1
	} else {
		dp[0][1] = 1
	}
	for i := 1; i < len(s); i++ {
		// dp[i][0]需要满足全为0
		// dp[i][1]只要满足最后一个元素为1就行，那么前i-1个元素既可以为0，也可以为1
		if s[i] == '1' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = Min(dp[i-1][0], dp[i-1][1])
		} else {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = Min(dp[i-1][0], dp[i-1][1]) + 1
		}
	}
	return Min(dp[len(dp)-1][0], dp[len(dp)-1][1])
}
