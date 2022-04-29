/*
-------------------------------------
# @Time    : 2022/4/29 12:23:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 905_按奇偶排序数组.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
