/*
-------------------------------------
# @Time    : 2022/7/10 20:48:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_068_I_二叉搜索树的最近公共祖先.go
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
	for root != nil {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}
