/*
-------------------------------------
# @Time    : 2021/11/18 0:02:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 559_N 叉树的最大深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	root := &Node559{1, []*Node559{{3, []*Node559{{5, []*Node559{}}, {6, []*Node559{}}}}, {2, []*Node559{}}, {4, []*Node559{}}}}
	fmt.Println(maxDepthN(root))
}

func maxDepthN(root *Node559) int {
	// 递归
	if root == nil {
		return 0
	}
	depth := 0
	for i := 0; i < len(root.Children); i++ {
		depth = max559(depth, maxDepthN(root.Children[i]))
	}
	return 1 + depth

	// 迭代
	//if root == nil {
	//	return 0
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//depth := 0
	//for queue.Len() > 0 {
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*Node559)
	//		for j := 0; j < len(node.Children); j++ {
	//			queue.PushBack(node.Children[j])
	//		}
	//	}
	//	depth++
	//}
	//return depth
}

type Node559 struct {
	Val      int
	Children []*Node559
}

func max559(x, y int) int {
	if x > y {
		return x
	}
	return y
}
