/*
-------------------------------------
# @Time    : 2021/11/24 9:55:05
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 098_验证二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	// 迭代
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true

	// 递归
	//return dfs(root, math.MinInt64, math.MaxInt64)
}

//func dfs(root *TreeNode, lower, upper int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val >= upper || root.Val <= lower {
//		return false
//	}
//	return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
//}
