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
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
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
