/*
-------------------------------------
# @Time    : 2021/12/11 19:56:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 037_解数独.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}

func solveSudoku(board [][]byte) {
	var isValid func(row, col int, val byte) bool
	isValid = func(row, col int, val byte) bool {
		for i := 0; i < 9; i++ {
			if board[row][i] == val {
				return false
			}
		}
		for j := 0; j < 9; j++ {
			if board[j][col] == val {
				return false
			}
		}
		startRow := row / 3 * 3
		startCol := col / 3 * 3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == val {
					return false
				}
			}
		}
		return true
	}
	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] != '.' {
					continue
				}
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k)) {
						board[i][j] = byte(k)
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtrack()
}
