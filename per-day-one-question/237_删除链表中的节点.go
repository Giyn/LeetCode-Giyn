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

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	listNode4 := ListNode{Val: 9}
	listNode3 := ListNode{Val: 1, Next: &listNode4}
	listNode2 := ListNode{Val: 5, Next: &listNode3}
	listNode1 := ListNode{Val: 4, Next: &listNode2}

	head := &listNode1
	deleteNode(&listNode2)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
