/*
-------------------------------------
# @Time    : 2022/4/28 0:14:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 300_最长递增子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) (ans int) {
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		ans = Max(ans, dp[i])
	}
	return
}
