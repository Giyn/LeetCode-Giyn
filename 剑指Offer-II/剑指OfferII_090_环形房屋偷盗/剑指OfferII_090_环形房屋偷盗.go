/*
-------------------------------------
# @Time    : 2022/3/20 15:11:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_090_环形房屋偷盗.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	_rob := func(start, end int) int {
		if start == end {
			return nums[start]
		}
		dp := make([]int, len(nums))
		dp[start] = nums[start]
		dp[start+1] = max(nums[start], nums[start+1])
		for i := start + 2; i <= end; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
		return dp[end]
	}
	return max(_rob(0, len(nums)-2), _rob(1, len(nums)-1))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
