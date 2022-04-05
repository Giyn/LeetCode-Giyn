/*
-------------------------------------
# @Time    : 2022/3/11 12:20:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 234_回文链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 2, 2, 1})
	fmt.Println(isPalindromeList(head))
}

func isPalindromeList(head *ListNode) bool {
	// 找到链表中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = slow.Next

	// 反转后半部分
	var pre *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}

	// 对比两部分
	for pre != nil && head != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}
