/*
-------------------------------------
# @Time    : 2021/11/18 1:02:48
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 222_完全二叉树的节点个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	fmt.Println(countNodes(root))
}

func countNodes(root *TreeNode) int {
	// 利用完全二叉树特性
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return (2 << leftH) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)

	// 递归
	//if root == nil {
	//	return 0
	//}
	//return 1 + countNodes(root.Left) + countNodes(root.Right)
}
