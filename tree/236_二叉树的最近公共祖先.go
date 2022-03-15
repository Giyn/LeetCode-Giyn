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
	root := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}
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
	// 左右两边都不为空，则根节点为祖先
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
