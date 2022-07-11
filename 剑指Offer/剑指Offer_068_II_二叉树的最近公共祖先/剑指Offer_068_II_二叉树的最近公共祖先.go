/*
-------------------------------------
# @Time    : 2022/7/8 23:18:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_068_II_二叉树的最近公共祖先.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{}, Right: &TreeNode{Val: 8}}}
	p := root.Left
	q := root.Right
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
