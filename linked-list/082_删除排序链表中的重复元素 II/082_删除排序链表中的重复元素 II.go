/*
-------------------------------------
# @Time    : 2022/4/25 17:22:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 082_删除排序链表中的重复元素 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 2, 3, 3, 4, 5})
	res := deleteDuplicates(head)
	PrintLinkedList(res)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
