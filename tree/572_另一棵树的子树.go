/*
-------------------------------------
# @Time    : 2021/11/17 22:34:26
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 572_另一棵树的子树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}}
	subRoot := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(isSubtree(root, subRoot))
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val == b.Val {
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	return false
}
