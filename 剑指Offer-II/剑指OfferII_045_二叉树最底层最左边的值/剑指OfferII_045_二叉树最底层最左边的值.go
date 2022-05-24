/*
-------------------------------------
# @Time    : 2022/3/13 15:16:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_045_二叉树最底层最左边的值.go
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
	root := NewTreeNode([]int{1, 2, 3, 4, -1, 5, 6, -1, -1, 7})
	fmt.Println(findBottomLeftValue(root))
}

func findBottomLeftValue(root *TreeNode) (ans int) {
	maxDepth := math.MinInt64
	var backtrack func(node *TreeNode, depth int)
	backtrack = func(node *TreeNode, depth int) {
		if node.Left == nil && node.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				ans = node.Val
			}
			return
		}
		if node.Left != nil {
			backtrack(node.Left, depth+1)
		}
		if node.Right != nil {
			backtrack(node.Right, depth+1)
		}
		return
	}
	backtrack(root, 0)
	return
}
