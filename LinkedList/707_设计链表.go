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

import "fmt"

type MyLinkedList struct {
	fake *Node
}

type Node struct {
	Val  int
	Next *Node
}

func main() {
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.PrintLinkedList()
	list.AddAtIndex(1, 2)
	list.PrintLinkedList()
	fmt.Println(list.Get(1))
	list.DeleteAtIndex(1)
	list.PrintLinkedList()
	fmt.Println(list.Get(1))
}

func Constructor() MyLinkedList {
	fake := &Node{-1, nil}
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
	node := &Node{val, cur.Next}
	cur.Next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.fake
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{val, nil}
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
	node := &Node{val, cur.Next}
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

func (this *MyLinkedList) PrintLinkedList() {
	head := this.fake.Next
	for head != nil {
		fmt.Print(head.Val)
		head = head.Next
	}
	fmt.Println()
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
