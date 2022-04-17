/*
-------------------------------------
# @Time    : 2022/4/17 2:04:44
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 102_二叉树的层序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
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
