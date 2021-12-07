/*
-------------------------------------
# @Time    : 2021/12/7 16:03:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 047_全排列 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

func permuteUnique(nums []int) (ans [][]int) {
	var path []int
	sort.Ints(nums)
	vis := make([]bool, len(nums))
	var backtrack func(int)
	backtrack = func(index int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// !vis[i-1] 对树层去重, 递归已回退到该层
			if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i-1]) {
				continue
			}
			path = append(path, nums[i])
			vis[i] = true
			backtrack(index + 1)
			vis[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return
}
