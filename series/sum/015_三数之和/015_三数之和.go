/*
-------------------------------------
# @Time    : 2022/3/30 13:22:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 015_三数之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 > 0 {
				right--
			} else if n1+n2+n3 < 0 {
				left++
			} else {
				ans = append(ans, []int{n1, n2, n3})
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			}
		}
	}
	return
}
