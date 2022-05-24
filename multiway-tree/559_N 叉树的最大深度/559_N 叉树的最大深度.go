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
	root := &Node{Val: 1, Children: []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(maxDepthN(root))
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepthN(root *Node) int {
	// 递归
	if root == nil {
		return 0
	}
	depth := 0
	for i := 0; i < len(root.Children); i++ {
		depth = max(depth, maxDepthN(root.Children[i]))
	}
	return 1 + depth

	//迭代
	//if root == nil {
	//	return 0
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//depth := 0
	//for queue.Len() > 0 {
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*Node)
	//		for j := 0; j < len(node.Children); j++ {
	//			queue.PushBack(node.Children[j])
	//		}
	//	}
	//	depth++
	//}
	//return depth
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
