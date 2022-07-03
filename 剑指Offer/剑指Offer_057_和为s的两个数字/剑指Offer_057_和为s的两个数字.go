/*
-------------------------------------
# @Time    : 2022/7/3 18:00:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_057_和为s的两个数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{10, 26, 30, 31, 47, 60}
	target := 40
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return []int{}
}
