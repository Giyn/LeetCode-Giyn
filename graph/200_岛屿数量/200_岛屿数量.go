/*
-------------------------------------
# @Time    : 2021/12/7 9:09:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 200_岛屿数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) (ans int) {
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])

	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2' // 标记访问过了
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
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
