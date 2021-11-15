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
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode094{1, nil, &TreeNode094{2, &TreeNode094{3, nil, nil}, nil}}
	fmt.Println(inorderTraversal(root))
}

func inorderTraversal(root *TreeNode094) (ans []int) {
	stack := list.New()
	cur := root
	for cur != nil || stack.Len() != 0 {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode094)
			ans = append(ans, cur.Val)
			cur = cur.Right
		}
	}
	return
	//var inorder func(node *TreeNode094)
	//inorder = func(node *TreeNode094) {
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

type TreeNode094 struct {
	Val   int
	Left  *TreeNode094
	Right *TreeNode094
}
