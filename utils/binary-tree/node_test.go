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
	head := NewTreeNode([]int{1, 3, 2, 5, 3, -1, 9})
	PrintBinaryTree(head)
}
