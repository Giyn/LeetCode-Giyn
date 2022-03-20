/*
-------------------------------------
# @Time    : 2022/1/31 22:06:27
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
	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	// 递归
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

	// 迭代
	//ans = &ListNode{}
	//cur := ans
	//for list1 != nil && list2 != nil {
	//	if list1.Val <= list2.Val {
	//		cur.Next = &ListNode{list1.Val, nil}
	//		list1 = list1.Next
	//	} else {
	//		cur.Next = &ListNode{list2.Val, nil}
	//		list2 = list2.Next
	//	}
	//	cur = cur.Next
	//}
	//if list1 != nil {
	//	cur.Next = list1
	//} else {
	//	cur.Next = list2
	//}
	//return ans.Next
}
