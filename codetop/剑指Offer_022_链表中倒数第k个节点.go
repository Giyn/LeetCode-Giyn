/*
-------------------------------------
# @Time    : 2022/4/8 21:09:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_022_链表中倒数第k个节点.go
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
	k := 2
	fmt.Println(getKthFromEnd(head, k).Val)
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
