/*
-------------------------------------
# @Time    : 2021/11/1 0:58:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 206_反转链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type ListNode206 struct {
	Val  int
	Next *ListNode206
}

func main() {
	listNode5 := ListNode206{5, nil}
	listNode4 := ListNode206{4, &listNode5}
	listNode3 := ListNode206{3, &listNode4}
	listNode2 := ListNode206{2, &listNode3}
	listNode1 := ListNode206{1, &listNode2}

	head := &listNode1
	res := reverseList(head)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func reverseList(head *ListNode206) *ListNode206 {
	var pre *ListNode206
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
