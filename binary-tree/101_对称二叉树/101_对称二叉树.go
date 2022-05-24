/*
-------------------------------------
# @Time    : 2021/11/17 16:45:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 101_对称二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
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
		outside := compare(left.Left, right.Right)
		inside := compare(left.Right, right.Left)

		return outside && inside
	}
	return compare(root.Left, root.Right)

	// 迭代
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
