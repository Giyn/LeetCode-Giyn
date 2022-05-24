/*
-------------------------------------
# @Time    : 2022/3/29 9:59:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1004_最大连续1的个数 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	fmt.Println(longestOnes(nums, k))
}

func longestOnes(nums []int, k int) (ans int) {
	left, cnt := 0, 0
	for right := range nums {
		if nums[right] != 1 {
			cnt++
		}
		for cnt > k {
			if nums[left] != 1 {
				cnt--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
