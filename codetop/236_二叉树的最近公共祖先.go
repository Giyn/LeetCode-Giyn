/*
-------------------------------------
# @Time    : 2021/11/26 0:08:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 236_二叉树的最近公共祖先.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
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
