/*
-------------------------------------
# @Time    : 2022/7/3 16:20:46
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 二叉树第K层节点的个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{1, 2, 2, 3, 3, 3, 3, 4, 4})
	k := 3
	fmt.Println(getNumberOfKLevel(root, k))
}

func getNumberOfKLevel(root *TreeNode, k int) int {
	if root == nil || k < 1 {
		return 0
	}
	if k == 1 {
		return 1
	}
	// DFS
	// root.left 第 k-1 层节点的个数 + root.right 第 k-1 层节点的个数
	numLeft := getNumberOfKLevel(root.Left, k-1)
	numRight := getNumberOfKLevel(root.Right, k-1)
	return numLeft + numRight

	// BFS
	//queue := []*TreeNode{root}
	//for k > 1 {
	//	length := len(queue)
	//	for i := 0; i < length; i++ {
	//		node := queue[0]
	//		queue = queue[1:]
	//		if node.Left != nil {
	//			queue = append(queue, node.Left)
	//		}
	//		if node.Right != nil {
	//			queue = append(queue, node.Right)
	//		}
	//	}
	//	k--
	//}
	//return len(queue)
}
