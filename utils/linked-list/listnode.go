/*
-------------------------------------
# @Time    : 2022/3/10 14:00:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : listnode.go
# @Software: GoLand
-------------------------------------
*/

package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) *ListNode {
	var head *ListNode
	cur := new(ListNode)
	head = cur
	for _, val := range vals {
		cur.Next = &ListNode{val, nil}
		cur = cur.Next
	}
	return head.Next
}
