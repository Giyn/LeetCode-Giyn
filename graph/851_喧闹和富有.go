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
	inDeg := make([]int, n)
	for _, r := range richer {
		g[r[0]] = append(g[r[0]], r[1])
		inDeg[r[1]]++
	}

	ans = make([]int, n)
	for i := range ans {
		ans[i] = i
	}
	q := make([]int, 0, n)
	for i, deg := range inDeg {
		if deg == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for _, y := range g[x] {
			if quiet[ans[x]] < quiet[ans[y]] {
				ans[y] = ans[x] // 更新 x 的邻居的答案
			}
			inDeg[y]--
			if inDeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return
}
