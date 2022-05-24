/*
-------------------------------------
# @Time    : 2022/3/19 22:17:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_084_含有重复元素集合的全排列.go
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
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// !vis[i-1] 对树层去重, 递归已回退到该层
			if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i-1]) {
				continue
			}
			vis[i] = true
			path = append(path, nums[i])
			dfs()
			vis[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return
}
