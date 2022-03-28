/*
-------------------------------------
# @Time    : 2022/3/28 21:00:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 785_判断二分图.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	graph := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	fmt.Println(isBipartite(graph))
}

func isBipartite(graph [][]int) bool {
	visited := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if visited[i] != 0 {
			continue
		}
		queue := []int{i}
		visited[i] = 1
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			curColor := visited[cur]
			neighborColor := -curColor
			for _, neighbor := range graph[cur] {
				if visited[neighbor] == 0 {
					visited[neighbor] = neighborColor
					queue = append(queue, neighbor)
				} else if visited[neighbor] != neighborColor {
					return false
				}
			}
		}
	}
	return true
}
