/*
-------------------------------------
# @Time    : 2021/12/15 10:20:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 851_喧闹和富有.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	richer := [][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}}
	quiet := []int{3, 2, 5, 4, 6, 1, 7, 0}
	fmt.Println(loudAndRich(richer, quiet))
}

func loudAndRich(richer [][]int, quiet []int) (ans []int) {
	n := len(quiet)
	g := make([][]int, n)
	for _, r := range richer {
		g[r[1]] = append(g[r[1]], r[0])
	}
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	var dfs func(int)
	dfs = func(x int) {
		if ans[x] != -1 {
			return
		}
		ans[x] = x
		for _, y := range g[x] {
			dfs(y)
			if quiet[ans[y]] < quiet[ans[x]] {
				ans[x] = ans[y]
			}
		}
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	return
}
