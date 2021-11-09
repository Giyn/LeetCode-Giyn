/*
-------------------------------------
# @Time    : 2021/11/3 22:19:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 018_四数之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] == target {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return
}
