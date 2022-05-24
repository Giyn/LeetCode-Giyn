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
	root := &Node{1, &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil}, &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil}
	connect(root)
	fmt.Println(root.Left.Next.Val)
}

func connect(root *Node) *Node {
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
	//	var tmp []*Node
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*Node)
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
