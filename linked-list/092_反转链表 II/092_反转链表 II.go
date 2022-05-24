/*
-------------------------------------
# @Time    : 2022/3/23 23:49:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 092_反转链表 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 4, 5, 6, 7})
	left := 3
	right := 5
	res := reverseBetween(head, left, right)
	PrintLinkedList(res)
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}
