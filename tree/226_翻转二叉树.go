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

import "container/list"

func main() {
	root := &TreeNode226{4, &TreeNode226{2, &TreeNode226{1, nil, nil}, &TreeNode226{3, nil, nil}}, &TreeNode226{7, &TreeNode226{6, nil, nil}, &TreeNode226{9, nil, nil}}}
	invertTree(root)
}

func invertTree(root *TreeNode226) *TreeNode226 {
	// BFS
	if root == nil {
		return root
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode226)
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
	//	node := stack.Remove(stack.Back()).(*TreeNode226)
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

type TreeNode226 struct {
	Val   int
	Left  *TreeNode226
	Right *TreeNode226
}
