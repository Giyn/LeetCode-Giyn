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
	. "LeetCodeGiyn/utils"
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}
	fmt.Println(minDiffInBST(root))
}

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt64
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			ans = min(ans, abs(node.Val-node.Left.Val))
			cur := node.Left.Right
			for cur != nil {
				ans = min(ans, abs(node.Val-cur.Val))
				cur = cur.Right
			}
			dfs(node.Left)
		}
		if node.Right != nil {
			ans = min(ans, abs(node.Val-node.Right.Val))
			cur := node.Right.Left
			for cur != nil {
				ans = min(ans, abs(node.Val-cur.Val))
				cur = cur.Left
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
