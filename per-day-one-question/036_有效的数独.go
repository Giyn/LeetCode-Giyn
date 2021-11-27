/*
-------------------------------------
# @Time    : 2021/11/4 11:03:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 036_有效的数独.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	recordRow := [9][9]bool{}
	recordCol := [9][9]bool{}
	recordBox := [9][9]bool{}
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			if recordRow[i][index] == true || recordCol[index][j] == true || recordBox[j/3+(i/3)*3][index] == true {
				return false
			}
			recordRow[i][index] = true
			recordCol[index][j] = true
			recordBox[j/3+(i/3)*3][index] = true
		}
	}
	return true
}
