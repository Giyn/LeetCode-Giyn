/*
-------------------------------------
# @Time    : 2022/5/24 0:14:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 965_单值二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{2, 2, 2, 5, 2})
	fmt.Println(isUnivalTree(root))
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur.Val != val {
				return false
			}
			cur = cur.Right
		}
	}
	return true
}
