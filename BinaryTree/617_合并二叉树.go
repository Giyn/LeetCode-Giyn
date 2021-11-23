/*
-------------------------------------
# @Time    : 2021/11/23 21:34:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 617_合并二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root1 := &TreeNode617{1, &TreeNode617{3, &TreeNode617{5, nil, nil}, nil}, &TreeNode617{2, nil, nil}}
	root2 := &TreeNode617{2, &TreeNode617{1, nil, &TreeNode617{4, nil, nil}}, &TreeNode617{3, nil, &TreeNode617{7, nil, nil}}}
	root := mergeTrees(root1, root2)
	fmt.Println(root.Val)
}

func mergeTrees(root1 *TreeNode617, root2 *TreeNode617) *TreeNode617 {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	} else {
		root1.Val += root2.Val
	}

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}

type TreeNode617 struct {
	Val   int
	Left  *TreeNode617
	Right *TreeNode617
}
