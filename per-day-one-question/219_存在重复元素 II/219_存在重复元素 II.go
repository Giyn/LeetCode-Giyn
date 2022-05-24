/*
-------------------------------------
# @Time    : 2022/1/19 0:31:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 219_存在重复元素 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]bool, k+1)
	for i := 0; i < len(nums); i++ {
		if i > k {
			mp[nums[i-k-1]] = false
		}
		if mp[nums[i]] {
			return true
		}
		mp[nums[i]] = true
	}
	return false
}
