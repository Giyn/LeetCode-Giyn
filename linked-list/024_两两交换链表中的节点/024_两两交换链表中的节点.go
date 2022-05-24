/*
-------------------------------------
# @Time    : 2021/11/1 2:26:19
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
	fake := &ListNode{Val: -1, Next: head}
	cur := fake
	for cur.Next != nil && cur.Next.Next != nil {
		tmpLeft := cur.Next
		tmpRight := cur.Next.Next
		cur.Next = tmpRight
		tmpLeft.Next = tmpRight.Next
		tmpRight.Next = tmpLeft

		cur = tmpLeft
	}
	return fake.Next
}
