/*
-------------------------------------
# @Time    : 2022/3/14 13:24:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_056_二叉搜索树中两个节点之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{8, 6, 10, 5, 7, 9, 11})
	k := 12
	fmt.Println(findTarget(root, k))
}

func findTarget(root *TreeNode, k int) bool {
	var stack []*TreeNode
	var vals []int
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, root.Val)
			root = root.Right
		}
	}
	left, right := 0, len(vals)-1
	for left < right {
		if vals[left]+vals[right] == k {
			return true
		} else if vals[left]+vals[right] < k {
			left++
		} else {
			right--
		}
	}
	return false
}
