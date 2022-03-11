/*
-------------------------------------
# @Time    : 2022/3/7 0:40:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_008_和大于等于 target 的最短子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"math"
)

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := math.MaxInt64
	left, right := 0, 0
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			ans = Min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
