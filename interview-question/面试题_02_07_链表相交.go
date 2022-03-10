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

import "fmt"

type ListNode0207 struct {
	Val  int
	Next *ListNode0207
}

func main() {
	same := &ListNode0207{8, &ListNode0207{4, &ListNode0207{5, nil}}}
	headA := &ListNode0207{4, &ListNode0207{1, same}}
	headB := &ListNode0207{5, &ListNode0207{0, &ListNode0207{1, same}}}
	fmt.Println(getIntersectionNode(headA, headB).Val)
}

func getIntersectionNode(headA, headB *ListNode0207) *ListNode0207 {
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
