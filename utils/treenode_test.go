/*
-------------------------------------
# @Time    : 2022/5/24 15:55:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : treenode_test.go
# @Software: GoLand
-------------------------------------
*/

package utils

import "testing"

func TestNewTreeNode(t *testing.T) {
	head := NewTreeNode([]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7})
	PrintBinaryTree(head)
}
