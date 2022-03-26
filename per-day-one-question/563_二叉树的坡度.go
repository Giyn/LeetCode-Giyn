/*
-------------------------------------
# @Time    : 2021/11/18 8:50:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 563_二叉树的坡度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println(findTilt(root))
}

func findTilt(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSlope := dfs(node.Left)
		rightSlope := dfs(node.Right)
		ans += Abs(leftSlope - rightSlope)
		return leftSlope + rightSlope + node.Val
	}
	dfs(root)
	return
}
