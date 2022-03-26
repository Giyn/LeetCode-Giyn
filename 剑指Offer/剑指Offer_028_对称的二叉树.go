/*
-------------------------------------
# @Time    : 2022/3/26 1:06:15
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_028_对称的二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(root))
}

func isSymmetric(root *TreeNode) bool {
	// 递归
	if root == nil {
		return true
	}
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		inside := compare(left.Right, right.Left)
		outside := compare(left.Left, right.Right)
		return inside && outside
	}
	return compare(root.Left, root.Right)

	// BFS
	//if root == nil {
	//	return true
	//}
	//var queue []*TreeNode
	//queue = append(queue, root.Left)
	//queue = append(queue, root.Right)
	//for len(queue) > 0 {
	//	left := queue[0]
	//	queue = queue[1:]
	//	right := queue[0]
	//	queue = queue[1:]
	//	if left == nil && right == nil {
	//		continue
	//	}
	//	if left == nil || right == nil || left.Val != right.Val {
	//		return false
	//	}
	//	queue = append(queue, left.Left)
	//	queue = append(queue, right.Right)
	//	queue = append(queue, left.Right)
	//	queue = append(queue, right.Left)
	//}
	//return true
}
