/*
-------------------------------------
# @Time    : 2021/10/1 11:09:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 437_路径总和 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type TreeNode437 struct {
	Val   int
	Left  *TreeNode437
	Right *TreeNode437
}

func main() {
	treeNode9 := TreeNode437{1, nil, nil}
	treeNode8 := TreeNode437{-2, nil, nil}
	treeNode7 := TreeNode437{3, nil, nil}
	treeNode6 := TreeNode437{11, nil, nil}
	treeNode5 := TreeNode437{2, nil, &treeNode9}
	treeNode4 := TreeNode437{3, &treeNode7, &treeNode8}
	treeNode3 := TreeNode437{-3, nil, &treeNode6}
	treeNode2 := TreeNode437{5, &treeNode4, &treeNode5}
	treeNode1 := TreeNode437{10, &treeNode2, &treeNode3}

	root := &treeNode1
	fmt.Println(pathSum(root, 8))
}

func rootSum(root *TreeNode437, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode437, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
