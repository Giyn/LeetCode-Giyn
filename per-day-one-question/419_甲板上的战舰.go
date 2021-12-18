/*
-------------------------------------
# @Time    : 2021/12/18 2:14:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 419_甲板上的战舰.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	fmt.Println(countBattleships(board))
}

func countBattleships(board [][]byte) (ans int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (board[i][j] == 'X') && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				ans++
			}
		}
	}
	return
}
