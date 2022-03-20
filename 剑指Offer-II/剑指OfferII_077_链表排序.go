/*
-------------------------------------
# @Time    : 2022/3/18 23:13:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_077_链表排序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	head := NewListNode([]int{-1, 5, 3, 4, 0})
	ans := sortList(head)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func sortList(head *ListNode) *ListNode {
	merge := func(head1, head2 *ListNode) *ListNode {
		dummy := &ListNode{}
		tmp, tmp1, tmp2 := dummy, head1, head2
		for tmp1 != nil && tmp2 != nil {
			if tmp1.Val <= tmp2.Val {
				tmp.Next = tmp1
				tmp1 = tmp1.Next
			} else {
				tmp.Next = tmp2
				tmp2 = tmp2.Next
			}
			tmp = tmp.Next
		}
		if tmp1 != nil {
			tmp.Next = tmp1
		} else if tmp2 != nil {
			tmp.Next = tmp2
		}
		return dummy.Next
	}
	if head == nil {
		return head
	}
	// 统计链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 初始化
	dummy := &ListNode{Next: head}
	// 每次将链表拆分成若干个长度为subLength的子链表
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev := dummy
		cur := dummy.Next // 记录拆分链表的位置
		// 如果还没拆分完
		for cur != nil {
			// 拆分链表1
			head1 := cur // 第一个链表的头
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 拆分链表2
			head2 := cur.Next // 第二个链表的头（链表1的尾部的下一个位置）
			cur.Next = nil    // 断开第一个链表和第二个链表
			cur = head2       // 第二个链表的头重新赋给cur
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 断开第二个链表的尾部的next
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummy.Next
}
