/*
-------------------------------------
# @Time    : 2021/11/18 11:50:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 110_平衡二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	if Abs(leftHeight-rightHeight) > 1 {
		return -1
	} else {
		return 1 + Max(leftHeight, rightHeight)
	}
}
