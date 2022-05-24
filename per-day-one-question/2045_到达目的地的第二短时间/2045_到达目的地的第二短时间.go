/*
-------------------------------------
# @Time    : 2022/1/24 11:18:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2045_到达目的地的第二短时间.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	edges := [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}
	time := 3
	change := 5
	fmt.Println(secondMinimum(n, edges, time, change))
}

func secondMinimum(n int, edges [][]int, time, change int) (ans int) {
	graph := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	// dist[i][0] 表示从 1 到 i 的最短路长度，dist[i][1] 表示从 1 到 i 的严格次短路长度
	dist := make([][2]int, n+1)
	dist[1][1] = math.MaxInt32
	for i := 2; i <= n; i++ {
		dist[i] = [2]int{math.MaxInt32, math.MaxInt32}
	}
	type pair struct{ x, d int }
	q := []pair{{1, 0}}
	for dist[n][1] == math.MaxInt32 {
		p := q[0]
		q = q[1:]
		for _, y := range graph[p.x] {
			d := p.d + 1
			if d < dist[y][0] {
				dist[y][0] = d
				q = append(q, pair{y, d})
			} else if dist[y][0] < d && d < dist[y][1] {
				dist[y][1] = d
				q = append(q, pair{y, d})
			}
		}
	}

	for i := 0; i < dist[n][1]; i++ {
		if ans%(change*2) >= change {
			ans += change*2 - ans%(change*2)
		}
		ans += time
	}
	return
}
