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
	. "LeetCodeGiyn/utils/binary-tree"
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(getMinimumDifference(root))
}

func getMinimumDifference(root *TreeNode) int {
	// 中序遍历
	ans := math.MaxInt64
	var pre *TreeNode
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil {
			ans = Min(ans, node.Val-pre.Val)
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return ans

	// DFS
	//ans := math.MaxInt64
	//var dfs func(node *TreeNode)
	//dfs = func(node *TreeNode) {
	//	if node == nil {
	//		return
	//	}
	//	if node.Left != nil {
	//		ans = Min(ans, node.Val-node.Left.Val)
	//		cur := node.Left.Right
	//		for cur != nil {
	//			ans = Min(ans, node.Val-cur.Val)
	//			cur = cur.Right
	//		}
	//		dfs(node.Left)
	//	}
	//	if node.Right != nil {
	//		ans = Min(ans, node.Right.Val-node.Val)
	//		cur := node.Right.Left
	//		for cur != nil {
	//			ans = Min(ans, cur.Val-node.Val)
	//			cur = cur.Left
	//		}
	//		dfs(node.Right)
	//	}
	//}
	//dfs(root)
	//return ans
}
