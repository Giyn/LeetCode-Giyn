/*
-------------------------------------
# @Time    : 2021/10/28 0:18:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 203_移除链表元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	listNode4 := ListNode{Val: 4}
	listNode3 := ListNode{Val: 3, Next: &listNode4}
	listNode2 := ListNode{Val: 2, Next: &listNode3}
	listNode1 := ListNode{Val: 1, Next: &listNode2}

	head := &listNode1

	res := removeElements(head, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	fake := new(ListNode)
	fake.Next = head
	cur := fake
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return fake.Next
}
