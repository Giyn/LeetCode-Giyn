/*
-------------------------------------
# @Time    : 2022/5/24 15:53:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : listnode.go
# @Software: GoLand
-------------------------------------
*/

package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) *ListNode {
	var head *ListNode
	cur := new(ListNode)
	head = cur
	for _, val := range vals {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return head.Next
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
