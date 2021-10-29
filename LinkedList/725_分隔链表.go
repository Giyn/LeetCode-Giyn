/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 725_分隔链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type ListNode725 struct {
	Val  int
	Next *ListNode725
}

func main() {
	listNode10 := ListNode725{10, nil}
	listNode9 := ListNode725{9, &listNode10}
	listNode8 := ListNode725{8, &listNode9}
	listNode7 := ListNode725{7, &listNode8}
	listNode6 := ListNode725{6, &listNode7}
	listNode5 := ListNode725{5, &listNode6}
	listNode4 := ListNode725{4, &listNode5}
	listNode3 := ListNode725{3, &listNode4}
	listNode2 := ListNode725{2, &listNode3}
	listNode1 := ListNode725{1, &listNode2}

	head := &listNode1

	res := splitListToParts(head, 3)
	for _, i := range res {
		for i != nil {
			fmt.Println(i.Val)
			i = i.Next
		}
		fmt.Println("---")
	}
}

func splitListToParts(head *ListNode725, k int) []*ListNode725 {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	base, rest := n/k, n%k
	parts := make([]*ListNode725, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := base
		if i < rest {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}
