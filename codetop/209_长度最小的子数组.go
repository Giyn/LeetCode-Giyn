/*
-------------------------------------
# @Time    : 2022/5/2 23:53:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 209_长度最小的子数组.go
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
	ans := math.MaxInt64
	sum := 0
	for left, right := 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target && left <= right {
			ans = Min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
