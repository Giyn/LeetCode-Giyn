/*
-------------------------------------
# @Time    : 2021/11/18 8:50:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 563_二叉树的坡度.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode563{1, &TreeNode563{2, nil, nil}, &TreeNode563{3, nil, nil}}
	fmt.Println(findTilt(root))
}

func findTilt(root *TreeNode563) (ans int) {
	var dfs func(node *TreeNode563) int
	dfs = func(node *TreeNode563) int {
		if node == nil {
			return 0
		}
		leftSlope := dfs(node.Left)
		rightSlope := dfs(node.Right)
		ans += abs563(leftSlope - rightSlope)
		return leftSlope + rightSlope + node.Val
	}
	dfs(root)
	return
}

type TreeNode563 struct {
	Val   int
	Left  *TreeNode563
	Right *TreeNode563
}

func abs563(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
