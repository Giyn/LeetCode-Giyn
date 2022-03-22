/*
-------------------------------------
# @Time    : 2021/11/12 9:51:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 145_二叉树的后序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	fmt.Println(postorderTraversal(root))
}

func postorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	Reverse(ans, 0, len(ans)-1)
	return

	//var postorder func(node *TreeNode)
	//postorder = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	postorder(node.Left)
	//	postorder(node.Right)
	//	ans = append(ans, node.Val)
	//}
	//postorder(root)
	//return
}
