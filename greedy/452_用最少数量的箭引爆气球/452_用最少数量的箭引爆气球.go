/*
-------------------------------------
# @Time    : 2021/12/21 17:06:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 452_用最少数量的箭引爆气球.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	points := [][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}
	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans = 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			ans++
		} else {
			points[i][1] = min(points[i][1], points[i-1][1])
		}
	}
	return
}
