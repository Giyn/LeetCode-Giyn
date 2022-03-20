/*
-------------------------------------
# @Time    : 2022/3/20 12:59:23
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
	var dirs = []struct{ x, y int }{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2' // 访问过
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx < m && nx >= 0 && ny < n && ny >= 0 {
				dfs(nx, ny)
			}
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
