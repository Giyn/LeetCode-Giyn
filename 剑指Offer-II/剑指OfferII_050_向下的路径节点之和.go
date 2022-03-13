/*
-------------------------------------
# @Time    : 2022/3/13 23:38:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_050_向下的路径节点之和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
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
		cur += int64(node.Val) // 从根节点到当前节点的总和
		ans += preSum[cur-int64(targetSum)]
		preSum[cur]++
		dfs(node.Left, cur)
		dfs(node.Right, cur)
		preSum[cur]-- // 回溯
		return
	}
	dfs(root, 0)
	return
}
