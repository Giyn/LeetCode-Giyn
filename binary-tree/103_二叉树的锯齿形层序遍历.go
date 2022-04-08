/*
-------------------------------------
# @Time    : 2022/4/8 21:12:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 103_二叉树的锯齿形层序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	isReverse := false
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if isReverse {
			Reverse(tmp, 0, len(tmp)-1)
		}
		isReverse = !isReverse
		ans = append(ans, tmp)
	}
	return
}
