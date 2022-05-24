/*
-------------------------------------
# @Time    : 2021/11/1 3:02:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 019_删除链表的倒数第 N 个结点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	n := 2
	res := removeNthFromEnd(head, n)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fake := &ListNode{Val: -1, Next: head}
	fast, slow := fake, fake
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return fake.Next
}
