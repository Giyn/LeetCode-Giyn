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
	. "LeetCodeGiyn/utils/binary-tree"
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
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
	reverse145(ans)
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

func reverse145(s []int) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
