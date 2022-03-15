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
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	p := root.Left
	q := root.Right
	res := lowestCommonAncestorBST(root, p, q)
	fmt.Println(res.Val)
}

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
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
	//	left := lowestCommonAncestorBST(root.Left, p, q)
	//	if left != nil {
	//		return left
	//	}
	//}
	//if root.Val < p.Val && root.Val < q.Val {
	//	right := lowestCommonAncestorBST(root.Right, p, q)
	//	if right != nil {
	//		return right
	//	}
	//}
	//return root
}
