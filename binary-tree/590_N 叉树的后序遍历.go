/*
-------------------------------------
# @Time    : 2021/11/17 10:52:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 590_N 叉树的后序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &Node590{1, []*Node590{{3, []*Node590{{5, []*Node590{}}, {6, []*Node590{}}}}, {2, []*Node590{}}, {4, []*Node590{}}}}
	fmt.Println(postorder(root))
}

func postorder(root *Node590) (ans []int) {
	// 迭代
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*Node590)
		ans = append(ans, node.Val)
		for i := 0; i < len(node.Children); i++ {
			stack.PushBack(node.Children[i])
		}
	}

	for i := 0; i < (len(ans)-1)/2+1; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return

	// 递归
	//if root == nil {
	//	return
	//}
	//var post func(node *Node590)
	//post = func(node *Node590) {
	//	for _, child := range node.Children {
	//		post(child)
	//	}
	//	ans = append(ans, node.Val)
	//}
	//post(root)
	//return
}

type Node590 struct {
	Val      int
	Children []*Node590
}
