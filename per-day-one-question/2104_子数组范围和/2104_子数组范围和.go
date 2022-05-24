/*
-------------------------------------
# @Time    : 2022/3/4 0:42:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2104_子数组范围和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{4, -2, -3, 4, 1}
	fmt.Println(subArrayRanges(nums))
}

func subArrayRanges(nums []int) (ans int64) {
	for i := 0; i < len(nums); i++ {
		maxV := nums[i]
		minV := nums[i]
		for j := i + 1; j < len(nums); j++ {
			maxV = max(maxV, nums[j])
			minV = min(minV, nums[j])
			ans += int64(maxV - minV)
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
