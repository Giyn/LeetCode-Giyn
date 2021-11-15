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
	//root := &Node117{1, &Node117{2, &Node117{4, nil, nil, nil}, &Node117{5, nil, nil, nil}, nil}, &Node117{3, nil, &Node117{7, nil, nil, nil}, nil}, nil}
	root := &Node117{3, &Node117{9, nil, nil, nil}, &Node117{20, &Node117{15, nil, nil, nil}, &Node117{7, nil, nil, nil}, nil}, nil}
	connect2(root)
	fmt.Println(root.Right.Left.Next.Val)
}

func connect2(root *Node117) *Node117 {
	start := root
	for start != nil {
		var nextStart, last *Node117
		handle := func(cur *Node117) {
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
	//} else {
	//	root.Next = nil
	//}
	//queue := list.New()
	//queue.PushBack(root)
	//for queue.Len() > 0 {
	//	var tmp []*Node117
	//	length := queue.Len()
	//	for i := 0; i < length; i++ {
	//		node := queue.Remove(queue.Front()).(*Node117)
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

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}
