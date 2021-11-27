/*
-------------------------------------
# @Time    : 2021/11/2 0:25:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 237_删除链表中的节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type ListNode237 struct {
	Val  int
	Next *ListNode237
}

func main() {
	listNode4 := ListNode237{9, nil}
	listNode3 := ListNode237{1, &listNode4}
	listNode2 := ListNode237{5, &listNode3}
	listNode1 := ListNode237{4, &listNode2}

	head := &listNode1
	deleteNode(&listNode2)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func deleteNode(node *ListNode237) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
