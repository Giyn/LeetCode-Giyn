/*
-------------------------------------
# @Time    : 2021/11/27 15:03:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 538_把二叉搜索树转换为累加树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"container/list"
)

func main() {
	root := NewTreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8})
	PrintBinaryTree(convertBST(root))
}

func convertBST(root *TreeNode) *TreeNode {
	// 迭代
	var pre int
	stack := list.New()
	cur := root
	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Right
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			cur.Val += pre
			pre = cur.Val
			cur = cur.Left
		}
	}
	return root

	// 递归
	//var pre int
	//var traversal func(cur *TreeNode)
	//traversal = func(cur *TreeNode) {
	//	if cur == nil {
	//		return
	//	}
	//	traversal(cur.Right)
	//	cur.Val += pre
	//	pre = cur.Val
	//	traversal(cur.Left)
	//}
	//traversal(root)
	//return root
}
