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

import "fmt"

func main() {
	root := &TreeNode572{3, &TreeNode572{4, &TreeNode572{1, nil, nil}, &TreeNode572{2, nil, nil}}, &TreeNode572{5, nil, nil}}
	subRoot := &TreeNode572{4, &TreeNode572{1, nil, nil}, &TreeNode572{2, nil, nil}}
	fmt.Println(isSubtree(root, subRoot))
}

func isSubtree(root *TreeNode572, subRoot *TreeNode572) bool {
	if root == nil {
		return false
	}
	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func check(a, b *TreeNode572) bool {
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

type TreeNode572 struct {
	Val   int
	Left  *TreeNode572
	Right *TreeNode572
}
