/*
-------------------------------------
# @Time    : 2021/11/15 21:44:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 117_填充每个节点的下一个右侧节点指针 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	root := &Node{3, &Node{9, nil, nil, nil}, &Node{20, &Node{15, nil, nil, nil}, &Node{7, nil, nil, nil}, nil}, nil}
	connect(root)
	fmt.Println(root.Right.Left.Next.Val)
}

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root

	//if root == nil {
	//	return nil
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
