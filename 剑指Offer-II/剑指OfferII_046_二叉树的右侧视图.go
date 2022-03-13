/*
-------------------------------------
# @Time    : 2022/3/13 16:35:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_046_二叉树的右侧视图.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 2, 3, -1, 5, -1, 4})
	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var right int
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			right = node.Val
		}
		ans = append(ans, right)
	}
	return
}
