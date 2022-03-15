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
	"container/list"
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
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
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
