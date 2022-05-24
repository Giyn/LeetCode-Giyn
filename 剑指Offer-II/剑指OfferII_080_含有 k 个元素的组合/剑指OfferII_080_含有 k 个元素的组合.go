/*
-------------------------------------
# @Time    : 2022/3/19 18:58:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_080_含有 k 个元素的组合.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

func combine(n int, k int) (ans [][]int) {
	var path []int
	var dfs func(index int)
	dfs = func(index int) {
		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		// 剪枝
		if len(path)+n-index+1 < k {
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return
}
