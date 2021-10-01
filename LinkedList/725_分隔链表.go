/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 725_分隔链表.go
# @Software: GoLand
-------------------------------------
*/

package splitListToParts

type ListNode struct {
	Val int
    Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	base, rest := n / k, n % k
	parts := make([]*ListNode, k)
	for i, curr := 0, head; i < k && curr != nil; i++ {
		parts[i] = curr
		partSize := base
		if i < rest {
			partSize++
		}
		for j := 1; j < partSize; j++ {
			curr = curr.Next
		}
		curr, curr.Next = curr.Next, nil
	}
	return parts
}
