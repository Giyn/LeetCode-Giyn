/*
-------------------------------------
# @Time    : 2021/11/26 16:55:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 669_修剪二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{3, &TreeNode{0, nil, &TreeNode{2, &TreeNode{1, nil, nil}, nil}}, &TreeNode{4, nil, nil}}
	low := 1
	high := 3
	res := trimBST(root, low, high)
	fmt.Println(res.Left.Val)
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
