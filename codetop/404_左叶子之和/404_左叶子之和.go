/*
-------------------------------------
# @Time    : 2022/4/25 22:54:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 404_左叶子之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Println(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Val
	}
	return ans + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
