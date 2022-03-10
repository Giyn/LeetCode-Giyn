/*
-------------------------------------
# @Time    : 2022/3/10 15:44:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 160_相交链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	intersection := NewListNode([]int{8, 4, 5})
	headA := NewListNode([]int{4, 1})
	headB := NewListNode([]int{5, 6, 1})
	headA.Next, headB.Next = intersection, intersection
	fmt.Println(getIntersectionNode(headA, headB).Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
