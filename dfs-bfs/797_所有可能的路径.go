/*
-------------------------------------
# @Time    : 2022/3/28 19:31:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 797_所有可能的路径.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	path := []int{0}
	var dfs func(x int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for _, y := range graph[x] {
			path = append(path, y)
			dfs(y)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return
}
