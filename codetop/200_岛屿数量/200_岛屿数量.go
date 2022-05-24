/*
-------------------------------------
# @Time    : 2022/4/8 11:21:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 200_岛屿数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2' // 已访问
		for _, dir := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nx, ny := x+dir[0], y+dir[1]
			dfs(nx, ny)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ans += 1
			}
		}
	}
	return
}
