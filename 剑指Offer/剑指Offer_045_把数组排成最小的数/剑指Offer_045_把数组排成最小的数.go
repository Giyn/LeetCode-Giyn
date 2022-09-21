/*
-------------------------------------
# @Time    : 2022/9/21 23:20:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_045_把数组排成最小的数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(nums))
}

func minNumber(nums []int) (ans string) {
	// 将整数数组按字符串形式排序
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})

	for i := 0; i < len(nums); i++ {
		ans += fmt.Sprintf("%d", nums[i])
	}
	return
}
