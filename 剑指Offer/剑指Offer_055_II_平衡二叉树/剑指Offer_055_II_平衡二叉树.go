/*
-------------------------------------
# @Time    : 2022/7/3 15:40:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_055_II_平衡二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(maxDepth(root.Left)-maxDepth(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
