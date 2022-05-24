/*
-------------------------------------
# @Time    : 2021/10/14 15:10:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 209_长度最小的子数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	target := 15
	nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	ans, sum, n := math.MaxInt32, 0, len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, 0

	for right < n {
		sum += nums[right]
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
