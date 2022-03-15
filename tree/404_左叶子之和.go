/*
-------------------------------------
# @Time    : 2021/11/20 15:12:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 404_左叶子之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{}
	//root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Left.Val
	}
	return ans + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
