/*
-------------------------------------
# @Time    : 2021/11/26 10:22:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 701_二叉搜索树中的插入操作.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode701{4, &TreeNode701{2, &TreeNode701{1, nil, nil}, &TreeNode701{3, nil, nil}}, &TreeNode701{7, nil, nil}}
	res := insertIntoBST(root, 5)
	fmt.Println(res.Val)
}

func insertIntoBST(root *TreeNode701, val int) *TreeNode701 {
	// 递归
	if root == nil {
		return &TreeNode701{val, nil, nil}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root

	// 迭代
	//if root == nil {
	//	return &TreeNode701{val, nil, nil}
	//}
	//var pre *TreeNode701
	//cur := root
	//for cur != nil {
	//	if cur.Val > val {
	//		pre = cur
	//		cur = cur.Left
	//	} else {
	//		pre = cur
	//		cur = cur.Right
	//	}
	//}
	//if pre.Val > val {
	//	pre.Left = &TreeNode701{val, nil, nil}
	//} else {
	//	pre.Right = &TreeNode701{val, nil, nil}
	//}
	//return root
}

type TreeNode701 struct {
	Val   int
	Left  *TreeNode701
	Right *TreeNode701
}
