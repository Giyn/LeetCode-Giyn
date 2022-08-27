/*
-------------------------------------
# @Time    : 2022/8/27 21:59:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 207_课程表.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	var result []int

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
