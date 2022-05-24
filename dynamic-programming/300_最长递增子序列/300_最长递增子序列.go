/*
-------------------------------------
# @Time    : 2021/10/8 17:32:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 300_最长递增子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) (ans int) {
	if len(nums) == 1 {
		return 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return
}
