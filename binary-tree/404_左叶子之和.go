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

import "fmt"

func main() {
	root := &TreeNode404{}
	//root := &TreeNode404{1, &TreeNode404{2, &TreeNode404{4, nil, nil}, &TreeNode404{5, nil, nil}}, &TreeNode404{3, nil, nil}}
	fmt.Println(sumOfLeftLeaves(root))
}

func sumOfLeftLeaves(root *TreeNode404) (ans int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		ans += root.Left.Val
	}
	return ans + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

type TreeNode404 struct {
	Val   int
	Left  *TreeNode404
	Right *TreeNode404
}
