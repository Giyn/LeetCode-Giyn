/*
-------------------------------------
# @Time    : 2022/3/23 23:17:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_024_反转链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	res := reverseList(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}
