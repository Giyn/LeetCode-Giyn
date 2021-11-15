/*
-------------------------------------
# @Time    : 2021/11/15 16:06:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 116_填充每个节点的下一个右侧节点指针.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	root := &Node116{1, &Node116{2, &Node116{4, nil, nil, nil}, &Node116{5, nil, nil, nil}, nil}, &Node116{3, &Node116{6, nil, nil, nil}, &Node116{7, nil, nil, nil}, nil}, nil}
	connect(root)
	fmt.Println(root.Left.Next.Val)
}

func connect(root *Node116) *Node116 {
	// O(1) space
	if root == nil {
		return root
	}
	for left := root; left.Left != nil; left = left.Left {
		for node := left; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root

	//if root == nil {
	//	return nil
	//} else {
	//	root.Next = nil
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//for queue.Len() > 0 {
	//	var tmp []*Node116
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*Node116)
	//		if node.Left != nil {
	//			queue.PushBack(node.Left)
	//		}
	//		if node.Right != nil {
	//			queue.PushBack(node.Right)
	//		}
	//		tmp = append(tmp, node)
	//	}
	//	for i := 0; i < len(tmp)-1; i++ {
	//		tmp[i].Next = tmp[i+1]
	//	}
	//}
	//return root
}

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}
