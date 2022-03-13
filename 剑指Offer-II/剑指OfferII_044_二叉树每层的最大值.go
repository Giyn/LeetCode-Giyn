/*
-------------------------------------
# @Time    : 2022/3/13 10:50:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_044_二叉树每层的最大值.go
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
	root := NewTreeNode([]int{1, 3, 2, 5, 3, -1, 9})
	fmt.Println(largestValues(root))
}

func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, Max(tmp))
	}
	return
}
