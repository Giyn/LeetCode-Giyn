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

import "fmt"

func main() {
	root := &TreeNode236{3, &TreeNode236{5, &TreeNode236{6, nil, nil}, &TreeNode236{2, &TreeNode236{7, nil, nil}, &TreeNode236{4, nil, nil}}}, &TreeNode236{1, &TreeNode236{0, nil, nil}, &TreeNode236{8, nil, nil}}}
	p := root.Left
	q := root.Right
	res := lowestCommonAncestor(root, p, q)
	fmt.Println(res.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode236) *TreeNode236 {
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

type TreeNode236 struct {
	Val   int
	Left  *TreeNode236
	Right *TreeNode236
}
