/*
-------------------------------------
# @Time    : 2022/3/10 14:36:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : listnode_test.go
# @Software: GoLand
-------------------------------------
*/

package utils

import (
	"fmt"
	"testing"
)

func TestNewListNode(t *testing.T) {
	head := NewListNode([]int{1, 2})
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
