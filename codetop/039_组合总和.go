/*
-------------------------------------
# @Time    : 2022/5/3 13:46:20
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
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) (ans [][]int) {
	var path []int
	sort.Ints(candidates)
	var dfs func(remain, idx int)
	dfs = func(remain, idx int) {
		if remain == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := idx; i < len(candidates) && remain >= candidates[i]; i++ {
			path = append(path, candidates[i])
			dfs(remain-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return
}
