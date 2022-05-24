/*
-------------------------------------
# @Time    : 2022/3/19 21:43:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_082_含有重复元素集合的组合.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
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
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return
}
