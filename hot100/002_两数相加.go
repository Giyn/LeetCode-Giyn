/*
-------------------------------------
# @Time    : 2021/11/24 11:09:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 002_两数相加.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode002{2, &ListNode002{4, &ListNode002{3, nil}}}
	l2 := &ListNode002{5, &ListNode002{6, &ListNode002{4, nil}}}
	res := addTwoNumbers(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode002, l2 *ListNode002) (head *ListNode002) {
	var tail *ListNode002
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode002{sum, nil}
			tail = head
		} else {
			tail.Next = &ListNode002{sum, nil}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode002{carry, nil}
	}
	return
}

type ListNode002 struct {
	Val  int
	Next *ListNode002
}
