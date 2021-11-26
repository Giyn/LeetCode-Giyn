/*
-------------------------------------
# @Time    : 2021/11/26 9:36:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 235_二叉搜索树的最近公共祖先.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode235{6, &TreeNode235{2, &TreeNode235{0, nil, nil}, &TreeNode235{4, &TreeNode235{3, nil, nil}, &TreeNode235{5, nil, nil}}}, &TreeNode235{8, &TreeNode235{7, nil, nil}, &TreeNode235{9, nil, nil}}}
	p := root.Left
	q := root.Right
	res := lowestCommonAncestorBST(root, p, q)
	fmt.Println(res.Val)
}

func lowestCommonAncestorBST(root, p, q *TreeNode235) *TreeNode235 {
	if root == p || root == q || root == nil {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestorBST(root.Left, p, q)
		if left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestorBST(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}

type TreeNode235 struct {
	Val   int
	Left  *TreeNode235
	Right *TreeNode235
}
