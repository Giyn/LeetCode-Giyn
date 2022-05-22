/*
-------------------------------------
# @Time    : 2022/5/23 0:29:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 675_为高尔夫比赛砍树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	forest := [][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}}
	fmt.Println(cutOffTree(forest))
}

func cutOffTree(forest [][]int) (ans int) {
	var trees []struct{ dis, x, y int }
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, struct{ dis, x, y int }{h, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })
	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []struct{ dis, x, y int }{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				if x, y := p.x+d[0], p.y+d[1]; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, struct{ dis, x, y int }{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}
	preX, preY := 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return
}
