/*
-------------------------------------
# @Time    : 2022/4/29 12:50:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 427_建立四叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}}
	fmt.Println(construct(grid))
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	m, n := len(grid), len(grid[0])
	preSum := make([][]int, m+1)
	preSum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		preSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + grid[i][j]
		}
	}

	var dfs func(x0, y0, x1, y1 int) *Node
	dfs = func(x0, y0, x1, y1 int) *Node {
		if diff := preSum[x1][y1] - preSum[x1][y0] - preSum[x0][y1] + preSum[x0][y0]; diff == 0 {
			return &Node{false, true, nil, nil, nil, nil}
		} else if diff == (x1-x0)*(y1-y0) {
			return &Node{true, true, nil, nil, nil, nil}
		}
		hx, hy := (x0+x1)/2, (y0+y1)/2
		return &Node{true, false,
			dfs(x0, y0, hx, hy),
			dfs(x0, hy, hx, y1),
			dfs(hx, y0, x1, hy),
			dfs(hx, hy, x1, y1)}
	}
	return dfs(0, 0, m, n)
}
