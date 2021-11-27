/*
-------------------------------------
# @Time    : 2021/11/1 8:47:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 142_环形链表 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type ListNode142 struct {
	Val  int
	Next *ListNode142
}

func main() {
	listNode2 := ListNode142{2, nil}
	listNode4 := ListNode142{-4, &listNode2}
	listNode3 := ListNode142{0, &listNode4}
	listNode2.Next = &listNode3
	listNode1 := ListNode142{3, &listNode2}
	head := &listNode1
	res := detectCycle(head)
	cnt := 0
	for head != res {
		head = head.Next
		cnt++
	}
	fmt.Println(cnt)
}

func detectCycle(head *ListNode142) *ListNode142 {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
	}
	return nil
}
