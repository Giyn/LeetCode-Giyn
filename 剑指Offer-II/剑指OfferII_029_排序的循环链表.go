/*
-------------------------------------
# @Time    : 2022/3/11 19:42:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_029_排序的循环链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewNode([]int{3, 4, 1})
	head.Next.Next.Next = head
	x := 2
	ans := insert(head, x)
	for i := 0; i < 4; i++ {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func insert(aNode *Node, x int) *Node {
	// nil 特判
	if aNode == nil {
		node := &Node{Val: x}
		node.Next = node
		return node
	}
	// 找到真正的起点
	cur := aNode
	next := cur.Next
	for cur.Val <= next.Val {
		cur = cur.Next
		next = next.Next
		if next == aNode {
			break
		}
	}
	// 找到插入位置
	realNode := next
	for next.Val < x {
		cur = cur.Next
		next = next.Next
		if next == realNode {
			break
		}
	}

	node := &Node{Val: x, Next: next}
	cur.Next = node
	return aNode
}
