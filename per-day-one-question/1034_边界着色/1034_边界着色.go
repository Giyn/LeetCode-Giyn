/*
-------------------------------------
# @Time    : 2021/12/7 0:41:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1034_边界着色.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	gird := [][]int{{1, 1, 2}, {1, 1, 1}, {1, 1, 1}}
	row := 1
	col := 1
	color := 2
	fmt.Println(colorBorder(gird, row, col, color))
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	originColor := grid[row][col]
	type point struct{ x, y int }
	var boards []point

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isBoard := false
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != originColor {
				isBoard = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBoard {
			boards = append(boards, point{x, y})
		}
	}
	dfs(row, col)

	for _, board := range boards {
		grid[board.x][board.y] = color
	}

	return grid
}
