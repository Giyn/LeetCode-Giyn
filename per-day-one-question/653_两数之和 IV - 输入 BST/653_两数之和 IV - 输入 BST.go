/*
-------------------------------------
# @Time    : 2022/3/21 0:19:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 653_两数之和 IV - 输入 BST.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{8, 6, 10, 5, 7, 9, 11})
	k := 12
	fmt.Println(findTarget(root, k))
}

func findTarget(root *TreeNode, k int) bool {
	mp := make(map[int]bool)
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if mp[k-node.Val] {
			return true
		}
		mp[node.Val] = true
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
