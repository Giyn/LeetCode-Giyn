/*
-------------------------------------
# @Time    : 2021/11/1 0:58:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 206_反转链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	res := reverseList(head)
	PrintLinkedList(res)
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre

	// 递归
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//newHead := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//return newHead
}
