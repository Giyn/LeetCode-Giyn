/*
-------------------------------------
# @Time    : 2022/4/8 11:46:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 002_两数相加.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := addTwoNumbers(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	carry := 0
	var cur *ListNode
	for l1 != nil || l2 != nil || carry != 0 {
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
			head = &ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}
	return
}
