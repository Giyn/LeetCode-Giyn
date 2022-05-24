/*
-------------------------------------
# @Time    : 2022/4/17 14:59:42
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 141_环形链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	head := NewListNode([]int{3, 2, 0})
	head.Next.Next.Next = &ListNode{Val: -4, Next: head.Next}
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
