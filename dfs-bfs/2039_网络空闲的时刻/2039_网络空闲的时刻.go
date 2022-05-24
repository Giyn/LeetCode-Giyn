/*
-------------------------------------
# @Time    : 2022/3/20 10:51:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2039_网络空闲的时刻.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{0, 1}, {1, 2}}
	patience := []int{0, 2, 1}
	fmt.Println(networkBecomesIdle(edges, patience))
}

func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	n := len(patience)
	g := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	vis[0] = true
	queue := []int{0}
	for dist := 1; queue != nil; dist++ {
		tmp := queue
		queue = nil
		for _, x := range tmp {
			for _, v := range g[x] {
				if vis[v] {
					continue
				}
				vis[v] = true
				queue = append(queue, v)
				ans = max(ans, patience[v]*((2*dist-1)/patience[v])+2*dist+1)
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
