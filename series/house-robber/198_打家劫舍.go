/*
-------------------------------------
# @Time    : 2022/3/20 15:10:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 198_打家劫舍.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob1(nums))
}

func rob1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}
