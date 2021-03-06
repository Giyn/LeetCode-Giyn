/*
-------------------------------------
# @Time    : 2022/4/27 16:16:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 024_两两交换链表中的节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4})
	res := swapPairs(head)
	PrintLinkedList(res)
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		left, right := cur.Next, cur.Next.Next
		cur.Next = right
		left.Next = right.Next
		right.Next = left
		cur = left
	}
	return dummy.Next
}
