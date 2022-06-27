/*
-------------------------------------
# @Time    : 2022/6/28 4:46:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 061_旋转链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import . "LeetCodeGiyn/utils"

func main() {
	head := NewListNode([]int{1, 2})
	k := 1
	res := rotateRight(head, k)
	PrintLinkedList(res)
}

func rotateRight(head *ListNode, k int) (ans *ListNode) {
	if head == nil || k == 0 {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	cur.Next = head // 成环
	broken := length - k%length
	cur = head
	// 找到新链表的最后一个结点
	for broken > 1 {
		cur = cur.Next
		broken--
	}
	ans = cur.Next
	cur.Next = nil
	return
}
