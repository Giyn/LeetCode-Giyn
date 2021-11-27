/*
-------------------------------------
# @Time    : 2021/11/26 16:55:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 669_修剪二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode669{3, &TreeNode669{0, nil, &TreeNode669{2, &TreeNode669{1, nil, nil}, nil}}, &TreeNode669{4, nil, nil}}
	low := 1
	high := 3
	res := trimBST(root, low, high)
	fmt.Println(res.Left.Val)
}

func trimBST(root *TreeNode669, low int, high int) *TreeNode669 {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

type TreeNode669 struct {
	Val   int
	Left  *TreeNode669
	Right *TreeNode669
}
