/*
-------------------------------------
# @Time    : 2022/1/4 22:54:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 913_猫和老鼠.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	graph := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}
	fmt.Println(catMouseGame(graph))
}

func catMouseGame(graph [][]int) int {
	n := len(graph)
	memo := make([][][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, n)
			for k := range memo[i][j] {
				memo[i][j][k] = -2
			}
		}
	}

	var minMax func(i, m, c int) int
	minMax = func(i, m, c int) int {
		if i == 2*n {
			return 0
		}
		if m == 0 {
			return -1
		}
		if c == m {
			return 1
		}
		if memo[i][m][c] != -2 {
			return memo[i][m][c]
		}
		var res int
		if i%2 == 1 {
			res = -1
			for _, nxt := range graph[c] {
				if nxt != 0 {
					r := minMax(i+1, m, nxt)
					if r > res {
						res = r
					}
					if r == 1 {
						break
					}
				}
			}
		} else {
			res = 1
			for _, nxt := range graph[m] {
				r := minMax(i+1, nxt, c)
				if r < res {
					res = r
				}
				if r == -1 {
					break
				}
			}
		}
		memo[i][m][c] = res
		return res
	}

	ans := minMax(0, 1, 2)
	if ans == -1 {
		return 1
	} else if ans == 1 {
		return 2
	}
	return 0
}
