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

import "fmt"

type ListNode019 struct {
	Val  int
	Next *ListNode019
}

func main() {
	listNode5 := ListNode019{5, nil}
	listNode4 := ListNode019{4, &listNode5}
	listNode3 := ListNode019{3, &listNode4}
	listNode2 := ListNode019{2, &listNode3}
	listNode1 := ListNode019{1, &listNode2}

	head := &listNode1
	n := 2
	res := removeNthFromEnd(head, n)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func removeNthFromEnd(head *ListNode019, n int) *ListNode019 {
	fake := &ListNode019{-1, head}
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
