/*
-------------------------------------
# @Time    : 2021/12/2 22:15:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 216_组合总和 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	k := 3
	n := 9
	fmt.Println(combinationSum3(k, n))
}

func combinationSum3(k int, n int) (ans [][]int) {
	var path []int
	var backtrack func(targetSum int, k int, start int)
	backtrack = func(targetSum int, k int, start int) {
		if len(path) == k {
			if targetSum == 0 {
				ans = append(ans, append([]int{}, path...))
			}
			return
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			backtrack(targetSum-i, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(n, k, 1)
	return
}
