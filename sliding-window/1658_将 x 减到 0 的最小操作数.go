/*
-------------------------------------
# @Time    : 2022/5/2 17:55:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1658_将 x 减到 0 的最小操作数.go
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
	nums := []int{1, 1, 4, 2, 3}
	x := 5
	fmt.Println(minOperations(nums, x))
}

func minOperations(nums []int, x int) (ans int) {
	sum := 0
	ans = math.MaxInt64
	for _, num := range nums {
		sum += num
	}
	target := sum - x
	left, right := 0, 0
	window := 0
	for right < len(nums) {
		window += nums[right]
		for window > target && left <= right {
			window -= nums[left]
			left++
		}
		if window == target {
			ans = Min(ans, len(nums)-(right-left+1))
		}
		right++
	}
	if ans == math.MaxInt64 {
		return -1
	}
	return
}
