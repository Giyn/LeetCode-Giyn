/*
-------------------------------------
# @Time    : 2022/3/19 21:56:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_083_没有重复元素集合的全排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) (ans [][]int) {
	var path []int
	sort.Ints(nums)
	visited := make(map[int]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			dfs()
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return
}
