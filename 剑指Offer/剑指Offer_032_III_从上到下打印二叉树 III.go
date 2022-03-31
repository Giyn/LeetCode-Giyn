/*
-------------------------------------
# @Time    : 2022/3/31 16:27:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_032_III_从上到下打印二叉树 III.go
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
	fmt.Println(levelOrder3(root))
}

func levelOrder3(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	even := true
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
		if !even {
			Reverse(tmp, 0, len(tmp)-1)
		}
		ans = append(ans, tmp)
		even = !even
	}
	return
}
