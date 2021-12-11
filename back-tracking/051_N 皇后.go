/*
-------------------------------------
# @Time    : 2021/12/11 16:22:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 051_N 皇后.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 4
	fmt.Println(solveNQueens(n))
}

func solveNQueens(n int) (ans [][]string) {
	var isValid func(int, int, [][]string) bool
	isValid = func(row, col int, chessBoard [][]string) bool {
		// 检查列
		for i := 0; i < row; i++ {
			if chessBoard[i][col] == "Q" {
				return false
			}
		}
		// 检查左上角是否有皇后
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if chessBoard[i][j] == "Q" {
				return false
			}
		}
		// 检查右上角是否有皇后
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if chessBoard[i][j] == "Q" {
				return false
			}
		}
		return true
	}

	var backtrack func(int, [][]string)
	backtrack = func(row int, chessBoard [][]string) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = strings.Join(chessBoard[i], "")
			}
			ans = append(ans, temp)
			return
		}
		for col := 0; col < n; col++ {
			if isValid(row, col, chessBoard) {
				chessBoard[row][col] = "Q"
				backtrack(row+1, chessBoard)
				chessBoard[row][col] = "."
			}
		}
	}

	var chessBoard = make([][]string, n)
	for i := 0; i < n; i++ {
		chessBoard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessBoard[i][j] = "."
		}
	}
	backtrack(0, chessBoard)
	return
}
