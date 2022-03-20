/*
-------------------------------------
# @Time    : 2021/11/27 15:33:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1038_把二叉搜索树转换为累加树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
)

func main() {
	root := NewTreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	PrintBinaryTree(bstToGst(root))
}

func bstToGst(root *TreeNode) *TreeNode {
	var pre int
	var traversal func(cur *TreeNode)
	traversal = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		traversal(cur.Right)
		cur.Val += pre
		pre = cur.Val
		traversal(cur.Left)
	}
	traversal(root)
	return root
}
