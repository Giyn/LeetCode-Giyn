/*
-------------------------------------
# @Time    : 2021/11/16 9:33:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 391_完美矩形.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	rectangles := [][]int{{1, 1, 3, 3}, {3, 1, 4, 2}, {3, 2, 4, 4}, {1, 3, 2, 4}, {2, 3, 3, 4}}
	fmt.Println(isRectangleCover(rectangles))
}

func isRectangleCover(rectangles [][]int) bool {
	x, y, a, b, s := 10001, 10001, -10001, -10001, 0
	mp := map[int]int{} // 统计各点数量
	for _, r := range rectangles {
		x, y, a, b = min391(x, r[0]), min391(y, r[1]), max391(a, r[2]), max391(b, r[3])
		s += (r[2] - r[0]) * (r[3] - r[1])
		mp[point(r[0], r[1])] += 1
		mp[point(r[0], r[3])] += 1
		mp[point(r[2], r[1])] += 1
		mp[point(r[2], r[3])] += 1
	}
	// 面积相等
	if s != (a-x)*(b-y) {
		return false
	}
	// 矩阵最外边的点
	points := map[int]bool{}
	points[point(x, y)] = true
	points[point(x, b)] = true
	points[point(a, y)] = true
	points[point(a, b)] = true

	for p := range points {
		// 最外边的点重复, 一定会重叠
		if v, err := mp[p]; !err || v > 1 {
			return false
		}
	}
	for p, v := range mp {
		if !points[p] {
			if v > 4 || v%2 == 1 {
				return false
			}
		}
	}
	return true
}

func min391(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max391(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func point(a, b int) int {
	return 10001*a + b
}
