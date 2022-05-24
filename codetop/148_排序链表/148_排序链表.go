/*
-------------------------------------
# @Time    : 2022/4/27 23:51:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 148_排序链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{-1, 5, 3, 4, 0})
	res := sortList(head)
	PrintLinkedList(res)
}

func sortList(head *ListNode) *ListNode {
	mergeList := func(list1, list2 *ListNode) *ListNode {
		dummy := &ListNode{}
		cur := dummy
		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				cur.Next = list1
				list1 = list1.Next
			} else {
				cur.Next = list2
				list2 = list2.Next
			}
			cur = cur.Next
		}
		if list1 != nil {
			cur.Next = list1
		}
		if list2 != nil {
			cur.Next = list2
		}
		return dummy.Next
	}
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	dummy := &ListNode{Next: head}
	for subLen := 1; subLen < length; subLen <<= 1 {
		prev := dummy
		cur := dummy.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLen && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = mergeList(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummy.Next
}
