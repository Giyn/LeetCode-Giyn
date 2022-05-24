/*
-------------------------------------
# @Time    : 2022/3/30 23:25:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 061_旋转链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	k := 2
	res := rotateRight(head, k)
	PrintLinkedList(res)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	cur := head
	length := 1
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	cur.Next = head // 成环
	broken := length - k%length
	pre := head
	for broken > 1 {
		pre = pre.Next
		broken--
	}
	cur = pre.Next
	pre.Next = nil
	return cur
}
