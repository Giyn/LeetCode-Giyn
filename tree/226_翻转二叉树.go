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
	"container/list"
)

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	invertTree(root)
}

func invertTree(root *TreeNode) *TreeNode {
	// BFS
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root

	// DFS
	//if root == nil {
	//	return root
	//}
	//stack := list.New()
	//stack.PushBack(root)
	//for stack.Len() > 0 {
	//	node := stack.Remove(stack.Back()).(*TreeNode)
	//	node.Left, node.Right = node.Right, node.Left
	//	if node.Left != nil {
	//		stack.PushBack(node.Left)
	//	}
	//	if node.Right != nil {
	//		stack.PushBack(node.Right)
	//	}
	//}
	//return root

	// 递归
	//if root == nil {
	//	return root
	//}
	//root.Left, root.Right = root.Right, root.Left
	//if root.Left != nil {
	//	invertTree(root.Left)
	//}
	//if root.Right != nil {
	//	invertTree(root.Right)
	//}
	//return root
}
