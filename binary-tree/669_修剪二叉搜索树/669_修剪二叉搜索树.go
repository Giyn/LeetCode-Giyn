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
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4}}
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
