/*
-------------------------------------
# @Time    : 2021/10/29 22:50:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 707_设计链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
)

type MyLinkedList struct {
	fake *ListNode
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list.Get(1))
}

func Constructor() MyLinkedList {
	fake := &ListNode{Val: -1}
	return MyLinkedList{fake}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.fake.Next
	if cur == nil {
		return -1
	}
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index != 0 {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	cur := this.fake
	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.fake
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &ListNode{Val: val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.fake
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index != 0 {
		return
	}
	node := &ListNode{Val: val, Next: cur.Next}
	cur.Next = node
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.fake
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index == 0 && cur.Next != nil {
		cur.Next = cur.Next.Next
	} else if index == 0 && cur.Next == nil {
		return
	}
}
