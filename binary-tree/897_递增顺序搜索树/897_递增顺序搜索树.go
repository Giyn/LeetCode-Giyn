/*
-------------------------------------
# @Time    : 2022/3/14 10:33:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 897_递增顺序搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	root := NewTreeNode([]int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9})
	PrintBinaryTree(increasingBST(root))
}

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		// 在中序遍历的过程中修改节点指向
		cur.Right = node
		node.Left = nil
		cur = node
		inorder(node.Right)
	}
	inorder(root)
	return dummy.Right
}
