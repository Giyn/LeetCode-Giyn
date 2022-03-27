/*
-------------------------------------
# @Time    : 2022/3/26 14:09:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 543_二叉树的直径.go
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
	root := NewTreeNode([]int{1, 2, 3, 4, 5})
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = Max(ans, left+right)
		return Max(left, right) + 1
	}
	dfs(root)
	return
}
