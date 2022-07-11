/*
-------------------------------------
# @Time    : 2021/11/26 9:36:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 235_二叉搜索树的最近公共祖先.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}
	p := root.Left
	q := root.Right
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 迭代
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root

	// 递归
	//if root == p || root == q || root == nil {
	//	return root
	//}
	//if root.Val > p.Val && root.Val > q.Val {
	//	left := lowestCommonAncestor(root.Left, p, q)
	//	if left != nil {
	//		return left
	//	}
	//}
	//if root.Val < p.Val && root.Val < q.Val {
	//	right := lowestCommonAncestor(root.Right, p, q)
	//	if right != nil {
	//		return right
	//	}
	//}
	//return root
}
