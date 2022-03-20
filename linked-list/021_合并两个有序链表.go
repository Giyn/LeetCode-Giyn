/*
-------------------------------------
# @Time    : 2022/3/19 15:07:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 021_合并两个有序链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

func main() {
	list1 := NewListNode([]int{1, 2, 4})
	list2 := NewListNode([]int{1, 3, 4})
	ans := mergeTwoLists(list1, list2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}
