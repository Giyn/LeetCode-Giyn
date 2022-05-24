/*
-------------------------------------
# @Time    : 2022/1/31 22:06:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 021_合并两个有序链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	list1 := NewListNode([]int{1, 2, 4})
	list2 := NewListNode([]int{1, 3, 4})
	res := mergeTwoLists(list1, list2)
	PrintLinkedList(res)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
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
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
