/*
-------------------------------------
# @Time    : 2021/11/12 9:54:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 094_二叉树的中序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return

	//var inorder func(node *TreeNode)
	//inorder = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	inorder(node.Left)
	//	ans = append(ans, node.Val)
	//	inorder(node.Right)
	//}
	//inorder(root)
	//return
}
