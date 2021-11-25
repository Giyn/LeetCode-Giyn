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
	"fmt"
	"math"
)

func main() {
	root := &TreeNode783{4, &TreeNode783{2, &TreeNode783{1, nil, nil}, &TreeNode783{3, nil, nil}}, &TreeNode783{6, nil, nil}}
	fmt.Println(getMinimumDifference783(root))
}

func getMinimumDifference783(root *TreeNode783) int {
	ans := math.MaxInt64
	var dfs func(node *TreeNode783)
	dfs = func(node *TreeNode783) {
		if node == nil {
			return
		}
		if node.Left != nil {
			ans = min783(ans, abs783(node.Val-node.Left.Val))
			cur := node.Left.Right
			for cur != nil {
				ans = min783(ans, abs783(node.Val-cur.Val))
				cur = cur.Right
			}
			dfs(node.Left)
		}
		if node.Right != nil {
			ans = min783(ans, abs783(node.Val-node.Right.Val))
			cur := node.Right.Left
			for cur != nil {
				ans = min783(ans, abs783(node.Val-cur.Val))
				cur = cur.Left
			}
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

type TreeNode783 struct {
	Val   int
	Left  *TreeNode783
	Right *TreeNode783
}

func min783(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs783(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
