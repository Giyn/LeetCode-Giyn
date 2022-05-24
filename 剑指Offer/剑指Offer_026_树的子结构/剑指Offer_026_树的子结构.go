/*
-------------------------------------
# @Time    : 2022/3/25 21:45:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_026_树的子结构.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	A := NewTreeNode([]int{3, 4, 5, 1, 2})
	B := NewTreeNode([]int{4, 1})
	fmt.Println(isSubStructure(A, B))
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var dfs func(A, B *TreeNode) bool
	dfs = func(A, B *TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return dfs(A.Left, B.Left) && dfs(A.Right, B.Right)
	}
	return (A != nil && B != nil) && (dfs(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}
