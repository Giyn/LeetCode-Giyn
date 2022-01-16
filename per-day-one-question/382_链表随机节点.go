/*
-------------------------------------
# @Time    : 2022/1/16 10:39:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 382_链表随机节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	head := &ListNode382{1, &ListNode382{2, &ListNode382{3, nil}}}
	obj := Constructor(head)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

type ListNode382 struct {
	Val  int
	Next *ListNode382
}

type Solution struct {
	listNode *ListNode382
	length   int
}

func Constructor(head *ListNode382) Solution {
	count := 0
	tmp := head
	for tmp != nil {
		count++
		tmp = tmp.Next
	}
	return Solution{head, count}
}

func (this *Solution) GetRandom() int {
	index := rand.Intn(this.length)
	tmp := this.listNode
	for index > 0 {
		tmp = tmp.Next
		index--
	}
	return tmp.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
