/*
-------------------------------------
# @Time    : 2022/4/26 20:48:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 130_被围绕的区域.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#' // 搜索到'O'
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEdge := i == 0 || i == m-1 || j == 0 || j == n-1
			if isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O' // 能搜索到的即未被围绕
			} else if board[i][j] == 'O' {
				board[i][j] = 'X' // 被围绕
			}
		}
	}
}
