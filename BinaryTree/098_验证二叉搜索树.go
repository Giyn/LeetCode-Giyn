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
	"fmt"
	"math"
)

func main() {
	//root := &TreeNode098{2, &TreeNode098{1, nil, nil}, &TreeNode098{3, nil, nil}}
	root := &TreeNode098{5, &TreeNode098{1, nil, nil}, &TreeNode098{4, &TreeNode098{3, nil, nil}, &TreeNode098{6, nil, nil}}}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode098) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode098, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val >= upper || root.Val <= lower {
		return false
	}
	return dfs(root.Left, lower, root.Val) && dfs(root.Right, root.Val, upper)
}

type TreeNode098 struct {
	Val   int
	Left  *TreeNode098
	Right *TreeNode098
}
