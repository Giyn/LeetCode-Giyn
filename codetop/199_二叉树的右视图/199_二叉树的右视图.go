/*
-------------------------------------
# @Time    : 2022/4/26 0:03:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 199_二叉树的右视图.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}
	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode) (ans []int) {
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}
	dfs(root, 1)
	return
}
