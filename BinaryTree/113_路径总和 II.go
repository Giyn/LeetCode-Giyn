/*
-------------------------------------
# @Time    : 2021/11/21 22:10:04
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 113_路径总和 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	root := &TreeNode113{5, &TreeNode113{4, &TreeNode113{11, &TreeNode113{7, nil, nil}, &TreeNode113{2, nil, nil}}, nil}, &TreeNode113{8, &TreeNode113{13, nil, nil}, &TreeNode113{4, &TreeNode113{5, nil, nil}, &TreeNode113{1, nil, nil}}}}
	targetSum := 22
	fmt.Println(pathSum(root, targetSum))
}

func pathSum(root *TreeNode113, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(*TreeNode113, int)
	dfs = func(node *TreeNode113, remain int) {
		if node == nil {
			return
		}
		remain -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && remain == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, remain)
		dfs(node.Right, remain)
	}
	dfs(root, targetSum)
	return
}

type TreeNode113 struct {
	Val   int
	Left  *TreeNode113
	Right *TreeNode113
}
