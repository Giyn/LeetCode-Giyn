/*
-------------------------------------
# @Time    : 2022/4/25 20:56:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 129_求根节点到叶节点数字之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{0, 1})
	fmt.Println(sumNumbers(root))
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, prevSum int) int
	dfs = func(node *TreeNode, prevSum int) int {
		if node == nil {
			return 0
		}
		sum := 10*prevSum + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)
	}
	return dfs(root, 0)
}
