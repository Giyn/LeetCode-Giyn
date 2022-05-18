/*
-------------------------------------
# @Time    : 2022/5/18 18:18:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 210_课程表 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourses, prerequisites))
}

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	edges := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		inDegree[info[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		ans = append(ans, u)
		for _, v := range edges[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return
}
