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

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root := NewTreeNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1})
	targetSum := 8
	fmt.Println(pathSum(root, targetSum))
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(node *TreeNode, cur int64)
	dfs = func(node *TreeNode, cur int64) {
		if node == nil {
			return
		}
		cur += int64(node.Val)
		ans += preSum[cur-int64(targetSum)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]--
		return
	}
	dfs(root, 0)
	return

	//var rootSum func(root *TreeNode, targetSum int) int
	//rootSum = func(root *TreeNode, targetSum int) (res int) {
	//	if root == nil {
	//		return
	//	}
	//	val := root.Val
	//	if val == targetSum {
	//		res++
	//	}
	//	res += rootSum(root.Left, targetSum-val)
	//	res += rootSum(root.Right, targetSum-val)
	//	return
	//}
	//if root == nil {
	//	return 0
	//}
	//res := rootSum(root, targetSum)
	//res += pathSum(root.Left, targetSum)
	//res += pathSum(root.Right, targetSum)
	//return res
}
