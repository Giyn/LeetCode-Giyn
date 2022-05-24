/*
-------------------------------------
# @Time    : 2022/3/11 11:43:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 876_链表的中间结点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(middleNode(head).Val)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next == nil {
		return slow
	}
	return slow.Next
}
