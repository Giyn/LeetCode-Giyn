/*
-------------------------------------
# @Time    : 2022/5/24 15:54:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : treenode.go
# @Software: GoLand
-------------------------------------
*/

package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(vals []int) (root *TreeNode) {
	tree := make([]*TreeNode, len(vals))
	for i := 0; i < len(vals); i++ {
		var node *TreeNode
		if vals[i] != -1 {
			node = &TreeNode{vals[i], nil, nil}
		}
		tree[i] = node
		if i == 0 {
			root = node
		}
	}
	for i := 0; i*2+2 < len(vals); i++ {
		if tree[i] != nil {
			tree[i].Left = tree[i*2+1]
			tree[i].Right = tree[i*2+2]
		}
	}
	return
}

func PrintBinaryTree(root *TreeNode) {
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	var res [][]int
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				tmp = append(tmp, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			} else {
				tmp = append(tmp, -1)
			}
		}
		res = append(res, tmp)
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(res[i][j], " ")
		}
		fmt.Println()
	}
}
