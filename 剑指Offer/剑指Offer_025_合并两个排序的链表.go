/*
-------------------------------------
# @Time    : 2022/3/25 20:10:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_025_合并两个排序的链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	l1, l2 := NewListNode([]int{1, 2, 4}), NewListNode([]int{1, 3, 4})
	ans := mergeTwoLists(l1, l2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
