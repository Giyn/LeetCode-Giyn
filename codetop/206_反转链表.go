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
