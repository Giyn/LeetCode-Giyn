/*
-------------------------------------
# @Time    : 2021/11/16 9:00:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 111_二叉树的最小深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	root := &TreeNode111{3, &TreeNode111{9, nil, nil}, &TreeNode111{20, &TreeNode111{15, nil, nil}, &TreeNode111{7, nil, nil}}}
	fmt.Println(minDepth(root))
}

func minDepth(root *TreeNode111) int {
	// 递归
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	return 1 + min111(minDepth(root.Left), minDepth(root.Right))

	// 迭代
	//if root == nil {
	//	return 0
	//}
	//depth := 0
	//queue := list.New()
	//queue.PushBack(root)
	//for queue.Len() > 0 {
	//	length := queue.Len()
	//	depth++
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*TreeNode111)
	//		if node.Left != nil {
	//			queue.PushBack(node.Left)
	//		}
	//		if node.Right != nil {
	//			queue.PushBack(node.Right)
	//		}
	//		if node.Left == nil && node.Right == nil {
	//			return depth
	//		}
	//	}
	//}
	//return depth
}

type TreeNode111 struct {
	Val   int
	Left  *TreeNode111
	Right *TreeNode111
}

func min111(x, y int) int {
	if x < y {
		return x
	}
	return y
}
