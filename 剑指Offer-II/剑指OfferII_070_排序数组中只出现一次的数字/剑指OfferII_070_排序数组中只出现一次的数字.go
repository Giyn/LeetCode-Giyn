/*
-------------------------------------
# @Time    : 2022/3/16 19:33:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_070_排序数组中只出现一次的数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			right = mid
		}
	}
	return nums[left]

	// if len(nums) == 1 {
	// 	return nums[0]
	// }
	// mid := len(nums) / 2
	// if mid%2 == 0 && nums[mid-1] == nums[mid] {
	// 	ans = singleNonDuplicate(nums[:mid-1])
	// } else if mid%2 == 0 && nums[mid+1] == nums[mid] {
	// 	ans = singleNonDuplicate(nums[mid+2:])
	// } else if mid%2 == 1 && nums[mid-1] == nums[mid] {
	// 	ans = singleNonDuplicate(nums[mid+1:])
	// } else if mid%2 == 1 && nums[mid+1] == nums[mid] {
	// 	ans = singleNonDuplicate(nums[:mid])
	// } else {
	// 	return nums[mid]
	// }
	// return
}
