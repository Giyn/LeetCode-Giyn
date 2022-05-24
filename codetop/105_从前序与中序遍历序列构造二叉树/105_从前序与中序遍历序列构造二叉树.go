/*
-------------------------------------
# @Time    : 2022/4/25 21:32:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 105_从前序与中序遍历序列构造二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	PrintBinaryTree(root)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	var idx int
	for i, v := range inorder {
		if v == preorder[0] {
			idx = i // 找到根节点
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}
