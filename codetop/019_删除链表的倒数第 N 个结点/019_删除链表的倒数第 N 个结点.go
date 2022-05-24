/*
-------------------------------------
# @Time    : 2022/4/26 0:20:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 019_删除链表的倒数第 N 个结点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	n := 2
	res := removeNthFromEnd(head, n)
	PrintLinkedList(res)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
