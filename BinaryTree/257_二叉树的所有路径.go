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
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode257{1, &TreeNode257{2, nil, &TreeNode257{5, nil, nil}}, &TreeNode257{3, nil, nil}}
	fmt.Println(binaryTreePaths(root))
}

func binaryTreePaths(root *TreeNode257) (ans []string) {
	var travel func(node *TreeNode257, s string)
	travel = func(node *TreeNode257, s string) {
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

type TreeNode257 struct {
	Val   int
	Left  *TreeNode257
	Right *TreeNode257
}
