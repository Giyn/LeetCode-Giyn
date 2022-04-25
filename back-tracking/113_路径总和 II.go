/*
-------------------------------------
# @Time    : 2022/4/25 12:10:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 113_路径总和 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}
	targetSum := 22
	fmt.Println(pathSum(root, targetSum))
}

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(node *TreeNode, remain int)
	dfs = func(node *TreeNode, remain int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		remain -= node.Val
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && remain == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, remain)
		dfs(node.Right, remain)
	}
	dfs(root, targetSum)
	return
}
