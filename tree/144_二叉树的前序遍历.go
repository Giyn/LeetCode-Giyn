/*
-------------------------------------
# @Time    : 2021/11/12 9:37:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 144_二叉树的前序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(preorderTraversal(root))
}

func preorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return
}

//func preorderTraversal(root *TreeNode) (ans []int) {
//	var preorder func(node *TreeNode)
//	preorder = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		ans = append(ans, node.Val)
//		preorder(node.Left)
//		preorder(node.Right)
//	}
//	preorder(root)
//	return
//}
