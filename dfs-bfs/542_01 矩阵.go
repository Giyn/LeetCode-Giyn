/*
-------------------------------------
# @Time    : 2022/3/28 21:50:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 542_01 矩阵.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println(updateMatrix(mat))
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	var queue [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			x, y := node[0]+dir[0], node[1]+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == -1 {
				mat[x][y] = mat[node[0]][node[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return mat
}
