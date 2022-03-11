/*
-------------------------------------
# @Time    : 2022/3/11 11:42:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 143_重排链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	reorderList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func reorderList(head *ListNode) {
	// 求链表中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转链表后半部分
	var pre *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}

	// 合并链表
	for head != nil && pre != nil {
		tmp1 := head.Next
		tmp2 := pre.Next
		head.Next = pre
		head = tmp1
		pre.Next = head
		pre = tmp2
	}
}
