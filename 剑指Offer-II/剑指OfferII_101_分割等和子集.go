/*
-------------------------------------
# @Time    : 2022/3/11 22:11:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_101_分割等和子集.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = Max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
