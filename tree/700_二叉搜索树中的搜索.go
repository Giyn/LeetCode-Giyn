/*
-------------------------------------
# @Time    : 2021/11/23 21:46:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 700_二叉搜索树中的搜索.go
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
	res := searchBST(root, 2)
	fmt.Println(res.Val)
}

func searchBST(root *TreeNode, val int) *TreeNode {
	// 迭代
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root

	// 递归
	//if root == nil {
	//	return root
	//}
	//if root.Val == val {
	//	return root
	//} else if root.Val > val {
	//	return searchBST(root.Left, val)
	//} else {
	//	return searchBST(root.Right, val)
	//}
}
