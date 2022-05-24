/*
-------------------------------------
# @Time    : 2022/3/26 14:43:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 124_二叉树中的最大路径和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
	"math"
)

func main() {
	root := NewTreeNode([]int{-10, 9, 20, -1, -1, 15, 7})
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) (ans int) {
	ans = math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		innerMax := left + node.Val + right
		ans = max(ans, innerMax)
		outputMax := node.Val + max(left, right)
		return max(outputMax, 0)
	}
	dfs(root)
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
