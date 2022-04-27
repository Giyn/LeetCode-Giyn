/*
-------------------------------------
# @Time    : 2022/4/26 16:02:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 031_下一个排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 从后往前找到第一个相邻升序对
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 找到nums[j:]里面第一个大于nums[i]的数
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 反转nums[j:]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
