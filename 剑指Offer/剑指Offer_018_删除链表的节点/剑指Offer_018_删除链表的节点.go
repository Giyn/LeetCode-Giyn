/*
-------------------------------------
# @Time    : 2022/3/25 19:59:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_018_删除链表的节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{4, 5, 1, 9})
	val := 5
	res := deleteNode(head, val)
	PrintLinkedList(res)
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			if cur.Next.Next != nil {
				cur.Next = cur.Next.Next
			} else {
				cur.Next = nil
			}
		}
		cur = cur.Next
	}
	return dummy.Next
}
