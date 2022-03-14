/*
-------------------------------------
# @Time    : 2022/3/14 13:22:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 173_二叉搜索树迭代器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{7, 3, 15, -1, -1, 9, 20})
	bst := Constructor173(root)
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

func Constructor173(root *TreeNode) BSTIterator {
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

func (this *BSTIterator) Next() (ans int) {
	ans = this.queue[0]
	this.queue = this.queue[1:]
	return
}

func (this *BSTIterator) HasNext() bool {
	if len(this.queue) > 0 {
		return true
	}
	return false
}
