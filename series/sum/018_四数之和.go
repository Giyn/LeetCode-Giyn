/*
-------------------------------------
# @Time    : 2022/3/30 13:39:15
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
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		if i > 0 && nums[i-1] == n1 {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			n2 := nums[j]
			if j > i+1 && nums[j-1] == n2 {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				n3, n4 := nums[left], nums[right]
				if n1+n2+n3+n4 > target {
					right--
				} else if n1+n2+n3+n4 < target {
					left++
				} else {
					ans = append(ans, []int{n1, n2, n3, n4})
					for left < right && nums[left] == n3 {
						left++
					}
					for left < right && nums[right] == n4 {
						right--
					}
				}
			}
		}
	}
	return
}
