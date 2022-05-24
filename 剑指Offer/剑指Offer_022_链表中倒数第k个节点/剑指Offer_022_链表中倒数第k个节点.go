/*
-------------------------------------
# @Time    : 2022/3/24 10:16:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_022_链表中倒数第k个节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	head := NewListNode([]int{1, 2, 3, 4, 5})
	k := 2
	res := getKthFromEnd(head, k)
	PrintLinkedList(res)
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	ans := head
	cur := head
	for cur.Next != nil && k-1 > 0 {
		cur = cur.Next
		k--
	}
	for cur.Next != nil {
		cur = cur.Next
		ans = ans.Next
	}
	return ans
}
