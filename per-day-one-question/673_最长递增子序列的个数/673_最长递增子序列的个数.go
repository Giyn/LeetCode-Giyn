/*
-------------------------------------
# @Time    : 2021/10/8 20:41:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 673_最长递增子序列的个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println(findNumberOfLIS(nums))
}

func findNumberOfLIS(nums []int) int {
	maxLen := 0
	ans := 0
	n := len(nums)
	dp := make([]int, n)  // 记录LIS长度
	cnt := make([]int, n) // 记录LIS个数
	for i, x := range nums {
		dp[i] = 1
		cnt[i] = 1
		for j, y := range nums[:i] {
			if x > y {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = cnt[i]
		} else if dp[i] == maxLen {
			ans += cnt[i]
		}
	}
	return ans
}
