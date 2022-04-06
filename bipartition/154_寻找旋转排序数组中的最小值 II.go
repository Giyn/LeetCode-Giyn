/*
-------------------------------------
# @Time    : 2022/4/6 20:27:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 154_寻找旋转排序数组中的最小值 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numbers := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin2(numbers))
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right-- // 无法确定在哪个范围，故缩小空间
		}
	}
	return nums[left]
}
