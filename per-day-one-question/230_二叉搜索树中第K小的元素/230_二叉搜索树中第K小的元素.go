/*
-------------------------------------
# @Time    : 2021/10/17 21:42:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 230_二叉搜索树中第K小的元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	node2 := TreeNode{Val: 2}
	node1 := TreeNode{Val: 1, Right: &node2}
	node4 := TreeNode{Val: 4}
	root := TreeNode{Val: 3, Left: &node1, Right: &node4}
	k := 1
	fmt.Println(kthSmallest(&root, k))
}

func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var LDR func(root *TreeNode)
	LDR = func(root *TreeNode) {
		if root != nil {
			LDR(root.Left)
			k--
			if k == 0 {
				ans = root.Val
				return
			}
			LDR(root.Right)
		}
	}
	LDR(root)
	return ans
}
