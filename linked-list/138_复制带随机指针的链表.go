/*
-------------------------------------
# @Time    : 2022/5/9 0:59:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 138_复制带随机指针的链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type NodeRandom struct {
	Val          int
	Next, Random *NodeRandom
}

func main() {
	node := &NodeRandom{2, nil, nil}
	node.Random = node
	head := &NodeRandom{1, node, node}
	ans := copyRandomList(head)
	for ans != nil {
		fmt.Println(ans.Val)
		if ans.Random != nil {
			fmt.Println("Random.Val = ", ans.Random.Val)
		} else {
			fmt.Println(nil)
		}
		ans = ans.Next
	}
}

func copyRandomList(head *NodeRandom) (ans *NodeRandom) {
	if head == nil {
		return head
	}
	for node := head; node.Next != nil; node = node.Next.Next {
		node.Next = &NodeRandom{Val: node.Val, Next: node.Next}
	}
	for node := head; node.Next != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	ans = head.Next
	for node := head; node.Next != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}
	return
}
