/*
-------------------------------------
# @Time    : 2022/1/12 23:32:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1036_逃离大迷宫.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	blocked := [][]int{{0, 1}, {1, 0}}
	source := []int{0, 0}
	target := []int{0, 2}
	fmt.Println(isEscapePossible(blocked, source, target))
}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	bound, block, max, dir := 1000000, map[int]bool{}, len(blocked)*(len(blocked)-1)/2, [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, b := range blocked {
		block[hash(b)] = true
	}
	bfs := func(start, end []int) bool {
		list := [][]int{start}
		explored := map[int]bool{}
		explored[hash(start)] = true
		for i := 0; i < len(list) && len(list) <= max; i++ {
			for _, d := range dir {
				point := make([]int, 2)
				point[0] = list[i][0] + d[0]
				point[1] = list[i][1] + d[1]
				h := hash(point)
				if point[0] >= 0 && point[0] < bound && point[1] >= 0 && point[1] < bound && !explored[h] && !block[h] {
					if point[0] == end[0] && point[1] == end[1] {
						return true
					}
					explored[h] = true
					list = append(list, point)
				}
			}
		}
		return len(list) > max
	}

	return bfs(source, target) && bfs(target, source)
}

func hash(point []int) int {
	return point[0]*1000000 + point[1]
}
