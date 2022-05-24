/*
-------------------------------------
# @Time    : 2022/3/14 14:01:37
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 653_两数之和 IV - 输入 BST.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
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
