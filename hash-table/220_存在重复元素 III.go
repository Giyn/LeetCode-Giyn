/*
-------------------------------------
# @Time    : 2022/3/15 14:25:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 220_存在重复元素 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	k := 2
	t := 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	getID := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1 // [-4,-3,-2,-1],t = 3
	}
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1) // t+1操作是为了确保差值小于等于t的数能够落到一个桶中,如[0,1,2,3],t = 3
		if _, ok := mp[id]; ok {
			return true
		}
		if y, ok := mp[id-1]; ok && Abs(x-y) <= t {
			return true
		}
		if y, ok := mp[id+1]; ok && Abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}
