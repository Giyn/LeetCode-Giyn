/*
-------------------------------------
# @Time    : 2022/4/27 16:38:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 958_二叉树的完全性检验.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 2, 3, 4, 5, -1, 7})
	fmt.Println(isCompleteTree(root))
}

func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	var isEmpty bool
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				isEmpty = true
			} else {
				if isEmpty {
					return false
				}
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
	}
	return true
}
