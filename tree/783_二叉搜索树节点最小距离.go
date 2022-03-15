/*
-------------------------------------
# @Time    : 2021/11/25 10:13:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 783_二叉搜索树节点最小距离.go
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
	fmt.Println(getMinimumDifference783(root))
}

func getMinimumDifference783(root *TreeNode) int {
	ans := math.MaxInt64
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			ans = Min(ans, Abs(node.Val-node.Left.Val))
			cur := node.Left.Right
			for cur != nil {
				ans = Min(ans, Abs(node.Val-cur.Val))
				cur = cur.Right
			}
			dfs(node.Left)
		}
		if node.Right != nil {
			ans = Min(ans, Abs(node.Val-node.Right.Val))
			cur := node.Right.Left
			for cur != nil {
				ans = Min(ans, Abs(node.Val-cur.Val))
				cur = cur.Left
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}
