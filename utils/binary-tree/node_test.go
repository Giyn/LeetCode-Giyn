/*
-------------------------------------
# @Time    : 2022/3/13 14:40:45
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : node_test.go
# @Software: GoLand
-------------------------------------
*/

package binary_tree

import (
	"testing"
)

func TestNewTreeNode(t *testing.T) {
	head := NewTreeNode([]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7})
	PrintBinaryTree(head)
}
