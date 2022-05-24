/*
-------------------------------------
# @Time    : 2022/4/29 22:28:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 079_单词搜索.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		if board[x][y] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		vis[x][y] = true
		defer func() { vis[x][y] = false }() // 回溯
		for _, dir := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			if nx, ny := x+dir[0], y+dir[1]; nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[nx][ny] {
				if dfs(nx, ny, idx+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
