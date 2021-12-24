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

type TreeNode230 struct {
	Val   int
	Left  *TreeNode230
	Right *TreeNode230
}

func main() {
	node2 := TreeNode230{2, nil, nil}
	node1 := TreeNode230{1, nil, &node2}
	node4 := TreeNode230{4, nil, nil}
	root := TreeNode230{3, &node1, &node4}
	k := 1
	fmt.Println(kthSmallest(&root, k))
}

func kthSmallest(root *TreeNode230, k int) int {
	var ans int
	var LDR func(root *TreeNode230)
	LDR = func(root *TreeNode230) {
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
