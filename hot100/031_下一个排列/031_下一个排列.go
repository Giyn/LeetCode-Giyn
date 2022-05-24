/*
-------------------------------------
# @Time    : 2022/3/1 19:39:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 031_下一个排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	left, right := -1, -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i
			break
		}
	}
	if left == -1 {
		sort.Ints(nums)
		return
	}
	for j := n - 1; j > left; j-- {
		if nums[j] > nums[left] {
			right = j
			break
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	for i, j := left+1, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
