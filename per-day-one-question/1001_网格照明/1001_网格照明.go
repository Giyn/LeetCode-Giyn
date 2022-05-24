/*
-------------------------------------
# @Time    : 2022/2/11 0:08:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1001_网格照明.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 5
	lamps := [][]int{{0, 0}, {4, 4}}
	queries := [][]int{{1, 1}, {1, 0}}
	fmt.Println(gridIllumination(n, lamps, queries))
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	rowCnts, colCnts, lrCnts, rlCnts, points := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}, map[int]bool{}
	for _, lamp := range lamps {
		x, y := lamp[0], lamp[1]
		p := x*n + y
		if !points[p] {
			points[p] = true
			rowCnts[x] += 1
			colCnts[y] += 1
			lrCnts[x+y] += 1
			rlCnts[x-y] += 1
		}
	}

	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		x, y := queries[i][0], queries[i][1]
		if rowCnts[x] > 0 || colCnts[y] > 0 || lrCnts[x+y] > 0 || rlCnts[x-y] > 0 {
			ans[i] = 1
			for _, dir := range [][]int{{0, 0}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
				nx, ny := x+dir[0], y+dir[1]
				if nx >= 0 && ny >= 0 && nx < n && ny < n {
					p := nx*n + ny
					if points[p] {
						points[p] = false
						rowCnts[nx] -= 1
						colCnts[ny] -= 1
						lrCnts[nx+ny] -= 1
						rlCnts[nx-ny] -= 1
					}
				}
			}
		}
	}
	return ans
}
