/*
-------------------------------------
# @Time    : 2022/2/22 11:17:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 053_最大子数组和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) (ans int) {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans = dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return
}
