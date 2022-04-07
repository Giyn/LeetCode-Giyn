/*
-------------------------------------
# @Time    : 2022/4/8 0:32:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_054_二叉搜索树的第k大节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{5, 3, 6, 2, 4, -1, -1, 1})
	k := 3
	fmt.Println(kthLargest(root, k))
}

func kthLargest(root *TreeNode, k int) (ans int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 求第k大，则倒序中序遍历
		dfs(node.Right)
		k--
		if k == 0 {
			ans = node.Val
			return
		}
		dfs(node.Left)
	}
	dfs(root)
	return
}
