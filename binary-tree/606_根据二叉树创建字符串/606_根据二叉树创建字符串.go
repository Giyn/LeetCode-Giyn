/*
-------------------------------------
# @Time    : 2022/3/19 16:49:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 606_根据二叉树创建字符串.go
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
	root := NewTreeNode([]int{1, 2, 3, 4})
	fmt.Println(tree2str(root))
}

func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
