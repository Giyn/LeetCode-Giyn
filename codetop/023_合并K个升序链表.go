/*
-------------------------------------
# @Time    : 2022/3/28 19:55:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 023_合并K个升序链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	lists := []*ListNode{NewListNode([]int{1, 4, 5}), NewListNode([]int{1, 3, 4}), NewListNode([]int{2, 6})}
	ans := mergeKLists(lists)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) >> 1
	return merge(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func merge(l1, l2 *ListNode) *ListNode {
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
