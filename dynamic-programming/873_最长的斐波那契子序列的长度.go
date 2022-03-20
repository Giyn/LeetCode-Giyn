/*
-------------------------------------
# @Time    : 2022/3/20 20:35:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 873_最长的斐波那契子序列的长度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(lenLongestFibSubseq(arr))
}

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	mp := make(map[int]int, n)
	for i := 0; i < n; i++ {
		mp[arr[i]] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			dp[i][j] = 2
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			diff := arr[j] - arr[i]
			if v, ok := mp[diff]; ok {
				if v < i {
					dp[i][j] = Max(dp[i][j], dp[v][i]+1)
				}
			}
			ans = Max(ans, dp[i][j])
		}
	}
	if ans > 2 {
		return
	} else {
		return 0
	}
}
