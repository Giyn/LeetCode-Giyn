/*
-------------------------------------
# @Time    : 2022/7/3 13:25:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_053_I_在排序数组中查找数字 I.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		right = left
		for right < len(nums) && nums[right] == target {
			right++
		}
		return right - left
	}
	return 0
}
