/*
-------------------------------------
# @Time    : 2022/5/17 4:38:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_021_调整数组顺序使奇数位于偶数前面.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if left < right && nums[left]%2 == 1 {
			left++
		}
		if left < right && nums[right]%2 == 0 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	return nums
}
