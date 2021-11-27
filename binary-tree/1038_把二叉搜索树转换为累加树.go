/*
-------------------------------------
# @Time    : 2021/11/27 15:33:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1038_把二叉搜索树转换为累加树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode1038{4, &TreeNode1038{1, &TreeNode1038{0, nil, nil}, &TreeNode1038{2, nil, &TreeNode1038{3, nil, nil}}}, &TreeNode1038{6, &TreeNode1038{5, nil, nil}, &TreeNode1038{7, nil, &TreeNode1038{8, nil, nil}}}}
	res := bstToGst(root)
	fmt.Println(res.Val)
}

func bstToGst(root *TreeNode1038) *TreeNode1038 {
	var pre int
	var traversal func(cur *TreeNode1038)
	traversal = func(cur *TreeNode1038) {
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

type TreeNode1038 struct {
	Val   int
	Left  *TreeNode1038
	Right *TreeNode1038
}
