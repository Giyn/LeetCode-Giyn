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
	"fmt"
)

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	fmt.Println(list.Get(1))
}

type MyLinkedList struct {
	dummy *Node
}

type Node struct {
	Val  int
	Next *Node
}

func Constructor() MyLinkedList {
	dummy := &Node{Val: -1}
	return MyLinkedList{dummy}
}

func (ml *MyLinkedList) Get(index int) int {
	cur := ml.dummy.Next
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

func (ml *MyLinkedList) AddAtHead(val int) {
	cur := ml.dummy
	node := &Node{Val: val, Next: cur.Next}
	cur.Next = node
}

func (ml *MyLinkedList) AddAtTail(val int) {
	cur := ml.dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{Val: val}
}

func (ml *MyLinkedList) AddAtIndex(index int, val int) {
	cur := ml.dummy
	for cur.Next != nil && index > 0 {
		cur = cur.Next
		index--
	}
	if index != 0 {
		return
	}
	node := &Node{Val: val, Next: cur.Next}
	cur.Next = node
}

func (ml *MyLinkedList) DeleteAtIndex(index int) {
	cur := ml.dummy
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
