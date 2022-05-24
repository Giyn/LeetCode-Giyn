/*
-------------------------------------
# @Time    : 2022/3/13 17:13:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_049_从根节点到叶节点的路径数字之和.go
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
		sum := prevSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return sum
		}
		return dfs(node.Left, sum) + dfs(node.Right, sum)

	}
	return dfs(root, 0)
}
