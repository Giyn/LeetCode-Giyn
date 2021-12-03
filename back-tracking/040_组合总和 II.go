/*
-------------------------------------
# @Time    : 2021/12/3 21:33:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 040_组合总和 II.go
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
	fmt.Println(combinationSum2(candidates, target))
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	var path []int
	sort.Ints(candidates)
	var backtrack func(target, index int)
	backtrack = func(target, index int) {
		if target == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := index; i < len(candidates) && target-candidates[i] >= 0; i++ {
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(target, 0)
	return
}
