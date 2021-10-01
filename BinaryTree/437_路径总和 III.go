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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rootSum(root *TreeNode, targetSum int) (res int) {
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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
