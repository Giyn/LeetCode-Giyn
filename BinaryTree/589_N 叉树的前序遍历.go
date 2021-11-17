/*
-------------------------------------
# @Time    : 2021/11/17 10:25:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 589_N 叉树的前序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &Node589{1, []*Node589{{3, []*Node589{{5, []*Node589{}}, {6, []*Node589{}}}}, {2, []*Node589{}}, {4, []*Node589{}}}}
	fmt.Println(preorder(root))
}

func preorder(root *Node589) (ans []int) {
	// 迭代
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*Node589)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
		ans = append(ans, node.Val)
	}
	return

	// 递归
	//if root == nil {
	//	return
	//}
	//var pre func(node *Node589)
	//pre = func(node *Node589) {
	//	ans = append(ans, node.Val)
	//	for _, child := range node.Children {
	//		pre(child)
	//	}
	//}
	//pre(root)
	//return
}

type Node589 struct {
	Val      int
	Children []*Node589
}
