/*
-------------------------------------
# @Time    : 2022/3/22 10:35:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_007_重建二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(root.Val)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func([]int, []int) *TreeNode
	dfs = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}
		rootVal := preorder[0]
		root := &TreeNode{Val: rootVal}
		if len(preorder) == 1 {
			return root
		}

		var splitIndex int
		for splitIndex = 0; splitIndex < len(inorder); splitIndex++ {
			if inorder[splitIndex] == rootVal {
				break
			}
		}
		leftInorder := inorder[:splitIndex]
		rightInorder := inorder[splitIndex+1:]

		preorder = preorder[1:]

		leftPreorder := preorder[:len(leftInorder)]
		rightPreorder := preorder[len(leftInorder):]

		root.Left = dfs(leftPreorder, leftInorder)
		root.Right = dfs(rightPreorder, rightInorder)

		return root
	}
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return dfs(preorder, inorder)
}
