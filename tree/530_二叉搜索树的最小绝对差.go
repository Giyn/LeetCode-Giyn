/*
-------------------------------------
# @Time    : 2021/11/25 9:22:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 530_二叉搜索树的最小绝对差.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode530{4, &TreeNode530{2, &TreeNode530{1, nil, nil}, &TreeNode530{3, nil, nil}}, &TreeNode530{6, nil, nil}}
	fmt.Println(getMinimumDifference(root))
}

func getMinimumDifference(root *TreeNode530) int {
	// 中序遍历
	ans := math.MaxInt64
	var pre *TreeNode530
	var traversal func(node *TreeNode530)
	traversal = func(node *TreeNode530) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil {
			ans = min530(ans, node.Val-pre.Val)
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return ans

	// DFS
	//ans := math.MaxInt64
	//var dfs func(node *TreeNode530)
	//dfs = func(node *TreeNode530) {
	//	if node == nil {
	//		return
	//	}
	//	if node.Left != nil {
	//		ans = min530(ans, node.Val-node.Left.Val)
	//		cur := node.Left.Right
	//		for cur != nil {
	//			ans = min530(ans, node.Val-cur.Val)
	//			cur = cur.Right
	//		}
	//		dfs(node.Left)
	//	}
	//	if node.Right != nil {
	//		ans = min530(ans, node.Right.Val-node.Val)
	//		cur := node.Right.Left
	//		for cur != nil {
	//			ans = min530(ans, cur.Val-node.Val)
	//			cur = cur.Left
	//		}
	//		dfs(node.Right)
	//	}
	//}
	//dfs(root)
	//return ans
}

type TreeNode530 struct {
	Val   int
	Left  *TreeNode530
	Right *TreeNode530
}

func min530(x, y int) int {
	if x < y {
		return x
	}
	return y
}
