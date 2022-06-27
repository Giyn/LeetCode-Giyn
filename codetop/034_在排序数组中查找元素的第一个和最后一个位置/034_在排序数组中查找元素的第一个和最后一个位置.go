/*
-------------------------------------
# @Time    : 2022/6/27 16:01:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 034_在排序数组中查找元素的第一个和最后一个位置.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < n && nums[left] == target {
		right = left
		for right < n && nums[right] == target {
			right++
		}
		return []int{left, right - 1}
	}
	return []int{-1, -1}
}
