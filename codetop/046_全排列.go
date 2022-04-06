/*
-------------------------------------
# @Time    : 2021/12/7 15:47:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 046_全排列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) (ans [][]int) {
	var path []int
	visited := make(map[int]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			visited[nums[i]] = true
			dfs()
			path = path[:len(path)-1]
			visited[nums[i]] = false
		}
	}
	dfs()
	return
}
