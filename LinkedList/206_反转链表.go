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

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	listNode5 := ListNode{5, nil}
	listNode4 := ListNode{4, &listNode5}
	listNode3 := ListNode{3, &listNode4}
	listNode2 := ListNode{2, &listNode3}
	listNode1 := ListNode{1, &listNode2}

	head := &listNode1
	res := reverseList(head)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
