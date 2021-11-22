/*
-------------------------------------
# @Time    : 2021/11/21 9:12:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 112_路径总和.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode112{5, &TreeNode112{4, &TreeNode112{11, &TreeNode112{7, nil, nil}, &TreeNode112{2, nil, nil}}, nil}, &TreeNode112{8, &TreeNode112{13, nil, nil}, &TreeNode112{4, nil, &TreeNode112{1, nil, nil}}}}
	targetSum := 22
	fmt.Println(hasPathSum(root, targetSum))
}

func hasPathSum(root *TreeNode112, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

type TreeNode112 struct {
	Val   int
	Left  *TreeNode112
	Right *TreeNode112
}
