/*
-------------------------------------
# @Time    : 2021/12/3 17:00:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 039_组合总和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 7, 6, 3, 5, 1}
	target := 9
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	var path []int
	sort.Ints(candidates)
	var backtrack func(target, index int)
	backtrack = func(target, index int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
		}
		// 剪枝(要先排序)
		for i := index; i < len(candidates) && target-candidates[i] >= 0; i++ {
			path = append(path, candidates[i])
			backtrack(target-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	backtrack(target, 0)
	return
}
