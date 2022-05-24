/*
-------------------------------------
# @Time    : 2022/1/29 10:50:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1765_地图中的最高点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	isWater := [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println(highestPeak(isWater))
}

func highestPeak(isWater [][]int) [][]int {
	type pair struct {
		x, y int
	}
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	var q []pair
	for i, row := range isWater {
		for j, water := range row {
			if water == 1 {
				ans[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}
	var dirs = []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, dir := range dirs {
			if x, y := p.x+dir.x, p.y+dir.y; 0 <= x && x < m && 0 <= y && y < n && ans[x][y] == -1 {
				ans[x][y] = ans[p.x][p.y] + 1
				q = append(q, pair{x, y})
			}
		}
	}
	return ans
}
