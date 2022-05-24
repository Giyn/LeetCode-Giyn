/*
-------------------------------------
# @Time    : 2021/11/18 21:53:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 257_二叉树的所有路径.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Println(binaryTreePaths(root))
}

func binaryTreePaths(root *TreeNode) (ans []string) {
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			ans = append(ans, v)
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		if node.Left != nil {
			travel(node.Left, s)
		}
		if node.Right != nil {
			travel(node.Right, s)
		}
	}
	travel(root, "")
	return
}
