/*
-------------------------------------
# @Time    : 2021/11/1 2:26:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 024_两两交换链表中的节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type ListNode024 struct {
	Val  int
	Next *ListNode024
}

func main() {
	listNode4 := ListNode024{4, nil}
	listNode3 := ListNode024{3, &listNode4}
	listNode2 := ListNode024{2, &listNode3}
	listNode1 := ListNode024{1, &listNode2}

	head := &listNode1
	res := swapPairs(head)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func swapPairs(head *ListNode024) *ListNode024 {
	fake := &ListNode024{-1, head}
	cur := fake
	for cur.Next != nil && cur.Next.Next != nil {
		tmpLeft := cur.Next
		tmpRight := cur.Next.Next
		cur.Next = tmpRight
		tmpLeft.Next = tmpRight.Next
		tmpRight.Next = tmpLeft

		cur = tmpLeft
	}
	return fake.Next
}
