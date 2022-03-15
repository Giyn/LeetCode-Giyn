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
	. "LeetCodeGiyn/utils/multiway-tree"
	"container/list"
	"fmt"
)

func main() {
	root := &Node{1, []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(preorder(root))
}

func preorder(root *Node) (ans []int) {
	// 迭代
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*Node)
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
	//var pre func(node *Node)
	//pre = func(node *Node) {
	//	ans = append(ans, node.Val)
	//	for _, child := range node.Children {
	//		pre(child)
	//	}
	//}
	//pre(root)
	//return
}
