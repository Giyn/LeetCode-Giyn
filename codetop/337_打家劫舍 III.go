/*
-------------------------------------
# @Time    : 2022/1/28 21:42:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 337_打家劫舍 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	fmt.Println(rob3(root))
}

func rob3(root *TreeNode) int {
	var robTree func(cur *TreeNode) [2]int
	// 后序遍历
	robTree = func(cur *TreeNode) [2]int {
		if cur == nil {
			return [2]int{0, 0}
		}
		left := robTree(cur.Left)
		right := robTree(cur.Right)
		yesRob := cur.Val + left[0] + right[0]
		noRob := Max(left[0], left[1]) + Max(right[0], right[1])
		return [2]int{noRob, yesRob}
	}
	ans := robTree(root)
	return Max(ans[0], ans[1])
}
