/*
-------------------------------------
# @Time    : 2022/5/16 15:27:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 面试题_04_06_后继者.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	p := &TreeNode{Val: 1}
	root := &TreeNode{Val: 2, Left: p, Right: &TreeNode{Val: 3}}
	fmt.Println(inorderSuccessor(root, p).Val)
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
	for node := root; node != nil; {
		if node.Val < p.Val {
			node = node.Right
		} else if node.Val > p.Val {
			ans, node = node, node.Left
		} else if node.Right != nil {
			node = node.Right
			for node.Left != nil {
				node = node.Left
			}
			return node
		} else {
			return
		}
	}
	return
}
