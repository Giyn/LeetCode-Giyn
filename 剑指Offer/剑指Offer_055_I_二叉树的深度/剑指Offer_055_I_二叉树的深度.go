/*
-------------------------------------
# @Time    : 2022/7/3 14:27:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_055_I_二叉树的深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{3, 9, 20 - 1, -1, 15, 7})
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
