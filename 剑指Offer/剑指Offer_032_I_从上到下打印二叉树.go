/*
-------------------------------------
# @Time    : 2022/3/31 16:14:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_032_I_从上到下打印二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(levelOrder1(root))
}

func levelOrder1(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			ans = append(ans, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return
}
