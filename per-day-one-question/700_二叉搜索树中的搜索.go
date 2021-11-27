/*
-------------------------------------
# @Time    : 2021/11/23 21:46:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 700_二叉搜索树中的搜索.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode700{4, &TreeNode700{2, &TreeNode700{1, nil, nil}, &TreeNode700{3, nil, nil}}, &TreeNode700{7, nil, nil}}
	res := searchBST(root, 2)
	fmt.Println(res.Val)
}

func searchBST(root *TreeNode700, val int) *TreeNode700 {
	// 迭代
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root

	// 递归
	//if root == nil {
	//	return root
	//}
	//if root.Val == val {
	//	return root
	//} else if root.Val > val {
	//	return searchBST(root.Left, val)
	//} else {
	//	return searchBST(root.Right, val)
	//}
}

type TreeNode700 struct {
	Val   int
	Left  *TreeNode700
	Right *TreeNode700
}
