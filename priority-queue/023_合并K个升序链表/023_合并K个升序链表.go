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
	. "LeetCodeGiyn/utils"
	"container/heap"
)

func main() {
	lists := []*ListNode{{1, &ListNode{Val: 4, Next: &ListNode{Val: 5}}}, {1, &ListNode{Val: 3, Next: &ListNode{Val: 4}}}, {2, &ListNode{Val: 6}}}
	res := mergeKLists(lists)
	PrintLinkedList(res)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := new(hp)
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

type hp []*ListNode

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
