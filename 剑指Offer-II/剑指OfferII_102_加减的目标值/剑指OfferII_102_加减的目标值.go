/*
-------------------------------------
# @Time    : 2022/3/21 9:34:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_102_加减的目标值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if abs(target) > sum || (target+sum)%2 == 1 {
		return 0
	}
	bag := (target + sum) / 2 // left - (sum - left) = target
	dp := make([]int, bag+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bag]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
