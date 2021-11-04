/*
-------------------------------------
# @Time    : 2021/11/5 0:28:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1218_最长定差子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	difference := 1
	fmt.Println(longestSubsequence(arr, difference))
}

func longestSubsequence(arr []int, difference int) (ans int) {
	dp := make(map[int]int)
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}
