/*
-------------------------------------
# @Time    : 2022/3/10 15:03:13
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_022_链表中环的入口节点.go
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
	fmt.Println(detectCycle(head).Val)
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
	}
	return nil
}
