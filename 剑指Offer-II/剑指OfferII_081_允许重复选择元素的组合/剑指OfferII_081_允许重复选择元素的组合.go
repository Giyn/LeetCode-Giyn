/*
-------------------------------------
# @Time    : 2022/3/19 21:18:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_081_允许重复选择元素的组合.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	var path []int
	sort.Ints(candidates)
	var dfs func(index, target int)
	dfs = func(index, target int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := index; i < len(candidates) && target-candidates[i] >= 0; i++ {
			path = append(path, candidates[i])
			dfs(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return
}
