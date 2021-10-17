/*
-------------------------------------
# @Time    : 2021/10/17 21:42:07
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 230_二叉搜索树中第K小的元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	node2 := TreeNode{2, nil, nil}
	node1 := TreeNode{1, nil, &node2}
	node4 := TreeNode{4, nil, nil}
	root := TreeNode{3, &node1, &node4}
	k := 1
	fmt.Println(kthSmallest(&root, k))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var LDR func(root *TreeNode)
	LDR = func(root *TreeNode) {
		if root != nil {
			LDR(root.Left)
			k--
			if k == 0 {
				ans = root.Val
				return
			}
			LDR(root.Right)
		}
	}
	LDR(root)
	return ans
}
