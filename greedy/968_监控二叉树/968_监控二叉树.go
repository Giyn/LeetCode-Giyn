/*
-------------------------------------
# @Time    : 2022/1/16 11:12:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 968_监控二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := &TreeNode{Left: &TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}}
	fmt.Println(minCameraCover(root))
}

func minCameraCover(root *TreeNode) (ans int) {
	// 0: 无覆盖, 1: 有摄像头, 2: 有覆盖
	var traversal func(cur *TreeNode) int
	traversal = func(cur *TreeNode) int {
		if cur == nil {
			return 2
		}
		left := traversal(cur.Left)
		right := traversal(cur.Right)

		// 情况1: 左右结点都有覆盖
		if left == 2 && right == 2 {
			return 0
		}

		// 情况2: 左右结点至少有一个无覆盖
		if left == 0 || right == 0 {
			ans++
			return 1
		}

		// 情况3: 左右结点至少有一个有摄像头
		if left == 1 || right == 1 {
			return 2
		}

		return -1 // 无意义逻辑
	}
	if traversal(root) == 0 {
		ans++
	}
	return
}
