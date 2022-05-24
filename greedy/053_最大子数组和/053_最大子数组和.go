/*
-------------------------------------
# @Time    : 2021/12/13 14:07:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 053_最大子数组和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) (ans int) {
	ans = math.MinInt64
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > ans {
			ans = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return
}
