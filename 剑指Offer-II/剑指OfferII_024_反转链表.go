/*
-------------------------------------
# @Time    : 2022/3/11 0:45:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_024_反转链表.go
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
	// 递归
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead

	// 迭代
	//var pre *ListNode
	//cur := head
	//for cur != nil {
	//	tmp := cur.Next
	//	cur.Next = pre
	//	pre = cur
	//	cur = tmp
	//}
	//return pre
}
