/*
-------------------------------------
# @Time    : 2021/11/9 11:26:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 面试题_02_07_链表相交.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	same := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: same}}
	headB := &ListNode{Val: 5, Next: &ListNode{Next: &ListNode{Val: 1, Next: same}}}
	fmt.Println(getIntersectionNode(headA, headB).Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}
