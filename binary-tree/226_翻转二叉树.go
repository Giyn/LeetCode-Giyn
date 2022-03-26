/*
-------------------------------------
# @Time    : 2021/11/16 21:07:22
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 226_翻转二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
)

func main() {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}}
	invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	// BFS
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root

	// DFS
	//if root == nil {
	//	return root
	//}
	//stack := []*TreeNode{root}
	//for len(stack) > 0 {
	//	node := stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	node.Left, node.Right = node.Right, node.Left
	//	if node.Left != nil {
	//		stack = append(stack, node.Left)
	//	}
	//	if node.Right != nil {
	//		stack = append(stack, node.Right)
	//	}
	//}
	//return root

	// 递归
	//if root == nil {
	//	return root
	//}
	//root.Left, root.Right = root.Right, root.Left
	//if root.Left != nil {
	//	mirrorTree(root.Left)
	//}
	//if root.Right != nil {
	//	mirrorTree(root.Right)
	//}
	//return root
}
