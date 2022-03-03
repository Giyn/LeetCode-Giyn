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
	"LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	nums := []int{4, -2, -3, 4, 1}
	fmt.Println(subArrayRanges(nums))
}

func subArrayRanges(nums []int) (ans int64) {
	for i := 0; i < len(nums); i++ {
		max := nums[i]
		min := nums[i]
		for j := i + 1; j < len(nums); j++ {
			max = utils.Max(max, nums[j])
			min = utils.Min(min, nums[j])
			ans += int64(max - min)
		}
	}
	return
}
