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

type Solution struct {
	listNode *ListNode
	//length   int
}

func Constructor382(head *ListNode) Solution {
	//count := 0
	//tmp := head
	//for tmp != nil {
	//	count++
	//	tmp = tmp.Next
	//}
	//return Solution{head, count}

	return Solution{head}
}

func (this *Solution) GetRandom() (ans int) {
	//index := rand.Intn(this.length)
	//tmp := this.listNode
	//for index > 0 {
	//	tmp = tmp.Next
	//	index--
	//}
	//return tmp.Val

	for node, index := this.listNode, 1; node != nil; node, index = node.Next, index+1 {
		if rand.Intn(index) == 0 {
			ans = node.Val
		}
	}
	return
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
