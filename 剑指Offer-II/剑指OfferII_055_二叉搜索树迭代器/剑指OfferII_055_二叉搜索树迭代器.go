/*
-------------------------------------
# @Time    : 2022/3/14 11:37:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_055_二叉搜索树迭代器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{7, 3, 15, -1, -1, 9, 20})
	bst := Constructor(root)
	fmt.Println(bst.Next())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
	fmt.Println(bst.Next())
	fmt.Println(bst.HasNext())
}

type BSTIterator struct {
	queue []int
}

func Constructor(root *TreeNode) BSTIterator {
	var queue []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		queue = append(queue, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return BSTIterator{queue}

	//var queue []int
	//var stack []*TreeNode
	//for root != nil || len(stack) != 0 {
	//	if root != nil {
	//		stack = append(stack, root)
	//		root = root.Left
	//	} else {
	//		root = stack[len(stack)-1]
	//		stack = stack[:len(stack)-1]
	//		queue = append(queue, root.Val)
	//		root = root.Right
	//	}
	//}
	//return BSTIterator{queue}
}

func (bst *BSTIterator) Next() (ans int) {
	ans = bst.queue[0]
	bst.queue = bst.queue[1:]
	return
}

func (bst *BSTIterator) HasNext() bool {
	if len(bst.queue) > 0 {
		return true
	}
	return false
}
