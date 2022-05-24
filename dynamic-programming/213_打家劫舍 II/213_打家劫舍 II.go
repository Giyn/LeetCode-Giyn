/*
-------------------------------------
# @Time    : 2022/1/28 17:46:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 213_打家劫舍 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	robRange := func(start, end int) int {
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
	n := len(nums)
	left := robRange(0, n-2)
	right := robRange(1, n-1)
	return max(left, right)
}
