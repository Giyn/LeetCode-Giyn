/*
-------------------------------------
# @Time    : 2022/3/31 16:23:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_032_II_从上到下打印二叉树 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
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
		ans = append(ans, tmp)
	}
	return
}
