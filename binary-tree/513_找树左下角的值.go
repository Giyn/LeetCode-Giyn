/*
-------------------------------------
# @Time    : 2021/11/20 16:52:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 513_找树左下角的值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, &TreeNode{5, &TreeNode{7, nil, nil}, nil}, &TreeNode{6, nil, nil}}}
	fmt.Println(findBottomLeftValue(root))
}

func findBottomLeftValue(root *TreeNode) (ans int) {
	// 递归
	maxDepth := math.MinInt64
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				ans = node.Val
			}
			return
		}
		if node.Left != nil {
			traversal(node.Left, depth+1)
		}
		if node.Right != nil {
			traversal(node.Right, depth+1)
		}
		return
	}
	traversal(root, 0)
	return

	// 迭代
	//queue := list.New()
	//queue.PushBack(root)
	//for queue.Len() > 0 {
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*TreeNode)
	//		if i == 0 {
	//			ans = node.Val
	//		}
	//		if node.Left != nil {
	//			queue.PushBack(node.Left)
	//		}
	//		if node.Right != nil {
	//			queue.PushBack(node.Right)
	//		}
	//	}
	//}
	//return
}
