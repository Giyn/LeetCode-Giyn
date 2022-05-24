/*
-------------------------------------
# @Time    : 2022/3/14 11:16:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_054_所有大于等于节点的值之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	root := NewTreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	PrintBinaryTree(convertBST(root))
}

func convertBST(root *TreeNode) *TreeNode {
	// 递归
	var pre int
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Right)
		node.Val += pre
		pre = node.Val
		postorder(node.Left)
	}
	postorder(root)
	return root

	// 遍历
	//var pre int
	//stack := list.New()
	//cur := root
	//for cur != nil || stack.Len() != 0 {
	//	if cur != nil {
	//		stack.PushBack(cur)
	//		cur = cur.Right
	//	} else {
	//		cur = stack.Remove(stack.Back()).(*TreeNode)
	//		cur.Val += pre
	//		pre = cur.Val
	//		cur = cur.Left
	//	}
	//}
	//return root
}
