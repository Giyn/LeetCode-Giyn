/*
-------------------------------------
# @Time    : 2021/11/26 10:22:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 701_二叉搜索树中的插入操作.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, nil, nil}}
	res := insertIntoBST(root, 5)
	fmt.Println(res.Val)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 递归
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root

	// 迭代
	//if root == nil {
	//	return &TreeNode{val, nil, nil}
	//}
	//var pre *TreeNode
	//cur := root
	//for cur != nil {
	//	if cur.Val > val {
	//		pre = cur
	//		cur = cur.Left
	//	} else {
	//		pre = cur
	//		cur = cur.Right
	//	}
	//}
	//if pre.Val > val {
	//	pre.Left = &TreeNode{val, nil, nil}
	//} else {
	//	pre.Right = &TreeNode{val, nil, nil}
	//}
	//return root
}
