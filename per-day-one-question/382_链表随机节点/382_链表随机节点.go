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
	. "LeetCodeGiyn/utils"
	"fmt"
	"math/rand"
)

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	obj := Constructor(head)
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}

type Solution struct {
	listNode *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

func (s *Solution) GetRandom() (ans int) {
	for node, index := s.listNode, 1; node != nil; node, index = node.Next, index+1 {
		if rand.Intn(index) == 0 {
			ans = node.Val
		}
	}
	return
}
