/*
-------------------------------------
# @Time    : 2022/3/26 11:01:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 141_环形链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{3, 2, 0})
	head.Next.Next.Next = &ListNode{Val: -4, Next: head.Next}
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
