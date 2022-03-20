/*
-------------------------------------
# @Time    : 2022/3/19 15:28:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 147_对链表进行插入排序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{-1, 5, 3, 4, 0})
	ans := insertionSortList(head)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	lastSorted := head
	cur := head.Next
	for cur != nil {
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}
			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}
	return dummy.Next
}
