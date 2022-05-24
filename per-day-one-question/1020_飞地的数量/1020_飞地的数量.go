/*
-------------------------------------
# @Time    : 2022/2/12 2:08:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1020_飞地的数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0},
	}
	fmt.Println(numEnclaves(grid))
}

func numEnclaves(grid [][]int) (ans int) {
	type pair struct {
		x, y int
	}
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var q []pair
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && grid[i][j] == 1 {
				vis[i][j] = true
				q = append(q, pair{i, j})
			}
		}
	}
	var dirs = []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, dir := range dirs {
			x, y := p.x+dir.x, p.y+dir.y
			if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 || vis[x][y] {
				continue
			}
			vis[x][y] = true
			q = append(q, pair{x, y})
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}
