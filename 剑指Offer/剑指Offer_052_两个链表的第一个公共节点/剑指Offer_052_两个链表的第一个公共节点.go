/*
-------------------------------------
# @Time    : 2022/7/2 18:33:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_052_两个链表的第一个公共节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	public := NewListNode([]int{8, 4, 5})
	headA := NewListNode([]int{4, 1})
	headB := NewListNode([]int{5, 6, 1})
	headA.Next, headB.Next = public, public
	fmt.Println(getIntersectionNode(headA, headB).Val)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
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
