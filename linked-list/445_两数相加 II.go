/*
-------------------------------------
# @Time    : 2022/3/11 10:52:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 445_两数相加 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	l1 := NewListNode([]int{6, 4, 5, 0, 4, 4, 9, 4, 1})
	l2 := NewListNode([]int{3, 8, 8, 0, 3, 0, 1, 4, 8})
	ans := addTwoNumbers2(l1, l2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1 []int
	var stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var carry int
	var head *ListNode
	for len(stack1) > 0 || len(stack2) > 0 {
		var v1, v2, sum int
		if len(stack1) > 0 {
			v1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			v2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum = (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		head = &ListNode{Val: sum, Next: head}
	}
	if carry > 0 {
		head = &ListNode{Val: carry, Next: head}
	}
	return head
}
