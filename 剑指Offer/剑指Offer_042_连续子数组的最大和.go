/*
-------------------------------------
# @Time    : 2022/5/20 21:48:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_042_连续子数组的最大和.go
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
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	sum := 0
	ans := math.MinInt64
	for _, num := range nums {
		sum += num
		ans = Max(ans, sum)
		if sum <= 0 {
			sum = 0
		}
	}
	return ans
}
