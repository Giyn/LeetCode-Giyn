/*
-------------------------------------
# @Time    : 2021/11/3 12:08:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 407_接雨水 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	heightMap := [][]int{{1, 4, 3, 1, 3, 2}, {3, 2, 1, 3, 2, 4}, {2, 3, 3, 2, 3, 1}}
	fmt.Println(trapRainWater(heightMap))
}

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return
	}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	h := &hp407{}
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{v, i, j})
				vis[i][j] = true
			}
		}
	}

	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && !vis[nx][ny] {
				if heightMap[nx][ny] < cur.h {
					ans += cur.h - heightMap[nx][ny]
				}
				vis[nx][ny] = true
				heap.Push(h, cell{max407(heightMap[nx][ny], cur.h), nx, ny})
			}
		}
	}
	return
}

type cell struct {
	h, x, y int
}
type hp407 []cell

func (h hp407) Len() int {
	return len(h)
}
func (h hp407) Less(i, j int) bool {
	return h[i].h < h[j].h
}
func (h hp407) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp407) Push(v interface{}) {
	*h = append(*h, v.(cell))
}
func (h *hp407) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func max407(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
