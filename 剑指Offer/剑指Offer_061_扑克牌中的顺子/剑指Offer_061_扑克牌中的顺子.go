/*
-------------------------------------
# @Time    : 2022/7/10 21:10:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_061_扑克牌中的顺子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 1, 2, 5}
	fmt.Println(isStraight(nums))
}

func isStraight(nums []int) bool {
	sort.Ints(nums)
	zero := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			zero++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[zero] < 5
}
