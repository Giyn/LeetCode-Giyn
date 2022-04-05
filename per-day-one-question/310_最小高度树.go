/*
-------------------------------------
# @Time    : 2022/4/6 0:21:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 310_最小高度树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
}

func findMinHeightTrees(n int, edges [][]int) (ans []int) {
	if n == 1 {
		return []int{0}
	}
	degree := make([]int, n)
	mp := make(map[int][]int)
	for _, edge := range edges {
		degree[edge[0]]++
		degree[edge[1]]++
		mp[edge[0]] = append(mp[edge[0]], edge[1])
		mp[edge[1]] = append(mp[edge[1]], edge[0])
	}
	var queue []int
	// 叶子结点入队
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		ans = []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			ans = append(ans, cur)
			neighbors := mp[cur]
			for _, neighbor := range neighbors {
				degree[neighbor]--
				if degree[neighbor] == 1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}
	return
}
