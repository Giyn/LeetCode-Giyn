/*
-------------------------------------
# @Time    : 2022/2/14 0:26:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 540_有序数组中的单一元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3}
	fmt.Println(singleNonDuplicate(nums))
}

func singleNonDuplicate(nums []int) (ans int) {
	if len(nums) == 1 {
		return nums[0]
	}
	mid := len(nums) / 2
	if mid%2 == 0 && nums[mid-1] == nums[mid] {
		ans = singleNonDuplicate(nums[:mid-1])
	} else if mid%2 == 0 && nums[mid+1] == nums[mid] {
		ans = singleNonDuplicate(nums[mid+2:])
	} else if mid%2 == 1 && nums[mid-1] == nums[mid] {
		ans = singleNonDuplicate(nums[mid+1:])
	} else if mid%2 == 1 && nums[mid+1] == nums[mid] {
		ans = singleNonDuplicate(nums[:mid])
	} else {
		return nums[mid]
	}
	return
}
