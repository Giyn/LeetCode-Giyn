/*
-------------------------------------
# @Time    : 2022/3/12 23:24:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 590_N 叉树的后序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/multiway-tree"
	"fmt"
)

func main() {
	root := &Node{Val: 1, Children: []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(postorder(root))
}

func postorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	var post func(node *Node)
	post = func(node *Node) {
		for _, child := range node.Children {
			post(child)
		}
		ans = append(ans, node.Val)
	}
	post(root)
	return
}
