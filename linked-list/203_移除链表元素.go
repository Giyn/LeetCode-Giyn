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

import "fmt"

type ListNode203 struct {
	Val  int
	Next *ListNode203
}

func main() {
	listNode4 := ListNode203{4, nil}
	listNode3 := ListNode203{3, &listNode4}
	listNode2 := ListNode203{2, &listNode3}
	listNode1 := ListNode203{1, &listNode2}

	head := &listNode1

	res := removeElements(head, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func removeElements(head *ListNode203, val int) *ListNode203 {
	fake := new(ListNode203)
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
