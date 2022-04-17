/*
-------------------------------------
# @Time    : 2022/4/17 15:07:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 025_K 个一组翻转链表.go
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
	ans := reverseKGroup(head, k)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	reverse := func(head *ListNode) *ListNode {
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
	dummy := &ListNode{Next: head}
	pre, end := dummy, dummy
	for end != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}
	return dummy.Next
}
