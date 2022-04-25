/*
-------------------------------------
# @Time    : 2022/4/25 22:32:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 144_二叉树的前序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) (ans []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
