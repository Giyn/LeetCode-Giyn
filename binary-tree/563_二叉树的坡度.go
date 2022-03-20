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
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
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
