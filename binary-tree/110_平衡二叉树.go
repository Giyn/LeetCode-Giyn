/*
-------------------------------------
# @Time    : 2021/11/18 11:50:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 110_平衡二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode110{3, &TreeNode110{9, nil, nil}, &TreeNode110{20, &TreeNode110{15, nil, nil}, &TreeNode110{7, nil, nil}}}
	fmt.Println(isBalanced(root))
}

func isBalanced(root *TreeNode110) bool {
	if root == nil {
		return true
	}
	if getHeight(root) == -1 {
		return false
	}
	return true
}

func getHeight(node *TreeNode110) int {
	if node == nil {
		return 0
	}
	leftHeight := getHeight(node.Left)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := getHeight(node.Right)
	if rightHeight == -1 {
		return -1
	}
	if abs110(leftHeight-rightHeight) > 1 {
		return -1
	} else {
		return 1 + max110(leftHeight, rightHeight)
	}
}

type TreeNode110 struct {
	Val   int
	Left  *TreeNode110
	Right *TreeNode110
}

func abs110(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func max110(x, y int) int {
	if x > y {
		return x
	}
	return y
}
