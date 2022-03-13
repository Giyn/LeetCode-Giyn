/*
-------------------------------------
# @Time    : 2022/3/13 16:41:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_047_二叉树剪枝.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
)

func main() {
	root := NewTreeNode([]int{1, 0, 1, 0, 0, 0, 1})
	PrintBinaryTree(pruneTree(root))
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
