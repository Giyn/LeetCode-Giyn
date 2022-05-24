/*
-------------------------------------
# @Time    : 2022/5/22 16:43:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 083_删除排序链表中的重复元素.go
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
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
