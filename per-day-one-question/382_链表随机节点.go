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
	. "LeetCodeGiyn/utils/linked-list"
	"fmt"
	"math/rand"
)

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	obj := Constructor382(head)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

type Solution382 struct {
	listNode *ListNode
}

func Constructor382(head *ListNode) Solution382 {
	return Solution382{head}
}

func (this *Solution382) GetRandom() (ans int) {
	for node, index := this.listNode, 1; node != nil; node, index = node.Next, index+1 {
		if rand.Intn(index) == 0 {
			ans = node.Val
		}
	}
	return
}
