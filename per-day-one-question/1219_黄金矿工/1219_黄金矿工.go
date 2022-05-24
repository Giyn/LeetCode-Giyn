/*
-------------------------------------
# @Time    : 2022/2/5 13:33:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1219_黄金矿工.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	fmt.Println(getMaximumGold(grid))
}

func getMaximumGold(grid [][]int) (ans int) {
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}
		rec := grid[x][y]
		grid[x][y] = 0
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[nx]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		grid[x][y] = rec
	}
	for i, row := range grid {
		for j, gold := range row {
			if gold > 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}
