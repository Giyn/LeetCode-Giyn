/*
-------------------------------------
# @Time    : 2021/11/27 15:03:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 538_把二叉搜索树转换为累加树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode538{4, &TreeNode538{1, &TreeNode538{0, nil, nil}, &TreeNode538{2, nil, &TreeNode538{3, nil, nil}}}, &TreeNode538{6, &TreeNode538{5, nil, nil}, &TreeNode538{7, nil, &TreeNode538{8, nil, nil}}}}
	res := convertBST(root)
	fmt.Println(res.Val)
}

func convertBST(root *TreeNode538) *TreeNode538 {
	var pre int
	var traversal func(cur *TreeNode538)
	traversal = func(cur *TreeNode538) {
		if cur == nil {
			return
		}
		traversal(cur.Right)
		cur.Val += pre
		pre = cur.Val
		traversal(cur.Left)
	}
	traversal(root)
	return root
}

type TreeNode538 struct {
	Val   int
	Left  *TreeNode538
	Right *TreeNode538
}
