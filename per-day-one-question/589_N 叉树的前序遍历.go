/*
-------------------------------------
# @Time    : 2022/3/11 0:08:35
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 589_N 叉树的前序遍历.go
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
	fmt.Println(preorder(root))
}

func preorder(root *Node) (ans []int) {
	if root == nil {
		return
	}
	var pre func(node *Node)
	pre = func(node *Node) {
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			pre(child)
		}
	}
	pre(root)
	return
}
