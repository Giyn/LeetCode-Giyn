/*
-------------------------------------
# @Time    : 2022/3/11 0:25:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 2049_统计最高分的节点数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	parents := []int{-1, 2, 0, 2, 0}
	fmt.Println(countHighestScoreNodes(parents))
}

func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	children := make([][]int, n)
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(node int) int {
		score, size := 1, n-1
		for _, ch := range children[node] {
			sz := dfs(ch)
			score *= sz
			size -= sz
		}
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		return n - size
	}
	dfs(0)
	return
}
