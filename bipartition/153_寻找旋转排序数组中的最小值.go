/*
-------------------------------------
# @Time    : 2022/4/6 17:04:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 153_寻找旋转排序数组中的最小值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin1(nums))
}

func findMin1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		// 选择无序的部分
		if nums[mid] < nums[right] {
			right = mid // 若mid-1，则丢失了最小值
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
