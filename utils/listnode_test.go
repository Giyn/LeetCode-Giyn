/*
-------------------------------------
# @Time    : 2022/5/24 15:55:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : listnode_test.go
# @Software: GoLand
-------------------------------------
*/

package utils

import "testing"

func TestNewListNode(t *testing.T) {
	head := NewListNode([]int{1, 2})
	PrintLinkedList(head)
}
