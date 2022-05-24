/*
-------------------------------------
# @Time    : 2022/3/30 17:00:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 016_最接近的三数之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 2}
	target := 3
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) (ans int) {
	sort.Ints(nums)
	ans = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum > target {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				return sum
			}
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
