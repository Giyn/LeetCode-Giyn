/*
-------------------------------------
# @Time    : 2021/11/29 10:03:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 023_合并K个升序链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/linked-list"
	"container/heap"
	"fmt"
)

func main() {
	lists := []*ListNode{{1, &ListNode{Val: 4, Next: &ListNode{Val: 5}}}, {1, &ListNode{Val: 3, Next: &ListNode{Val: 4}}}, {2, &ListNode{Val: 6}}}
	res := mergeKLists(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := new(hp023)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	dummyHead := new(ListNode)
	pre := dummyHead
	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		pre.Next = tmp
		pre = pre.Next
	}
	return dummyHead.Next
}

type hp023 []*ListNode

func (h hp023) Len() int            { return len(h) }
func (h hp023) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h hp023) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp023) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *hp023) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
