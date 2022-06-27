/*
-------------------------------------
# @Time    : 2022/6/28 2:19:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 662_二叉树最大宽度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 3, 2, 5, 3, -1, 9})
	fmt.Println(widthOfBinaryTree(root))
}

type Pair struct {
	idx  int
	node *TreeNode
}

func widthOfBinaryTree(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	queue := []Pair{{1, root}}
	for len(queue) > 0 {
		length := len(queue)
		ans = max(ans, queue[len(queue)-1].idx-queue[0].idx+1)
		for i := 0; i < length; i++ {
			item := queue[0]
			queue = queue[1:]
			if item.node.Left != nil {
				queue = append(queue, Pair{item.idx * 2, item.node.Left})
			}
			if item.node.Right != nil {
				queue = append(queue, Pair{item.idx*2 + 1, item.node.Right})
			}
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
