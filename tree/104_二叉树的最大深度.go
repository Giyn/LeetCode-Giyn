/*
-------------------------------------
# @Time    : 2021/11/15 23:38:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 104_二叉树的最大深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode) (ans int) {
	// 前序
	var getDepth func(node *TreeNode, depth int)
	getDepth = func(node *TreeNode, depth int) {
		if depth > ans {
			ans = depth
		}
		if node.Left == nil && node.Right == nil {
			return
		}
		if node.Left != nil {
			depth++
			getDepth(node.Left, depth)
			depth--
		}
		if node.Right != nil {
			depth++
			getDepth(node.Right, depth)
			depth--
		}
		return
	}
	if root == nil {
		return
	}
	getDepth(root, 1)
	return

	// 递归
	//if root == nil {
	//	return 0
	//}
	//return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))

	// 层序迭代
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
	//		node := queue.Remove(queue.Front()).(*TreeNode)
	//		if node.Left != nil {
	//			queue.PushBack(node.Left)
	//		}
	//		if node.Right != nil {
	//			queue.PushBack(node.Right)
	//		}
	//	}
	//}
	//return depth
}
