/*
-------------------------------------
# @Time    : 2022/3/14 10:34:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_053_二叉搜索树中的中序后继.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{5, 3, 6, 2, 4, -1, -1, 1})
	p := root.Left.Left
	fmt.Println(inorderSuccessor(root, p).Val)
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
	// 遍历
	cur := root
	for cur != nil {
		if cur.Val > p.Val {
			ans = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return

	// 递归
	//var pre *TreeNode
	//var dfs func(node *TreeNode)
	//dfs = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	dfs(node.Left)
	//	if pre == p {
	//		ans = node
	//	}
	//	pre = node
	//	dfs(node.Right)
	//}
	//dfs(root)
	//return
}
