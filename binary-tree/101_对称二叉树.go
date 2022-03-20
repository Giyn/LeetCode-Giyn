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
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
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
	//queue := list.New()
	//queue.PushBack(root.Left)
	//queue.PushBack(root.Right)
	//for queue.Len() > 0 {
	//	left := queue.Remove(queue.Front()).(*TreeNode)
	//	right := queue.Remove(queue.Front()).(*TreeNode)
	//	if left == nil && right == nil {
	//		continue
	//	}
	//	if left == nil || right == nil || left.Val != right.Val {
	//		return false
	//	}
	//	queue.PushBack(left.Left)
	//	queue.PushBack(right.Right)
	//	queue.PushBack(left.Right)
	//	queue.PushBack(right.Left)
	//}
	//return true
}
