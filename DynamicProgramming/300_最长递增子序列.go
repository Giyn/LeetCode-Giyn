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
	//nums := []int{0, 1, 0, 3, 2,3}
	//nums := []int{7, 7,7, 7, 7, 7, 7}
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i := range nums {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		ans = max(ans, dp[i])
	}

	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
