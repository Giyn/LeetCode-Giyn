/*
-------------------------------------
# @Time    : 2021/11/22 11:39:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 106_从中序与后序遍历序列构造二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTreeInPost(inorder, postorder)
	fmt.Println(root.Val)
}

func buildTreeInPost(inorder []int, postorder []int) *TreeNode {
	var dfs func([]int, []int) *TreeNode
	dfs = func(inorder []int, postorder []int) *TreeNode {
		if len(postorder) == 0 {
			return nil
		}
		rootVal := postorder[len(postorder)-1]
		root := &TreeNode{Val: rootVal}
		// 叶子结点
		if len(postorder) == 1 {
			return root
		}

		// 找切割点
		var splitIndex int
		for splitIndex = 0; splitIndex < len(inorder); splitIndex++ {
			if inorder[splitIndex] == rootVal {
				break
			}
		}

		// 切割中序数组 左闭右开
		leftInorder := inorder[:splitIndex]
		rightInorder := inorder[splitIndex+1:]

		postorder = postorder[:len(postorder)-1]

		// 切割后序数组 左闭右开
		leftPostorder := postorder[:len(leftInorder)]
		rightPostorder := postorder[len(leftInorder):]

		root.Left = dfs(leftInorder, leftPostorder)
		root.Right = dfs(rightInorder, rightPostorder)

		return root
	}
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	return dfs(inorder, postorder)
}
