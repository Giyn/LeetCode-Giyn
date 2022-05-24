/*
-------------------------------------
# @Time    : 2022/5/1 23:42:56
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1305_两棵二叉搜索树中的所有元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	root1 := NewTreeNode([]int{2, 1, 4})
	root2 := NewTreeNode([]int{1, 0, 3})
	fmt.Println(getAllElements(root1, root2))
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	inorder := func(root *TreeNode) (ans []int) {
		var dfs func(node *TreeNode)
		dfs = func(node *TreeNode) {
			if node == nil {
				return
			}
			dfs(node.Left)
			ans = append(ans, node.Val)
			dfs(node.Right)
			return
		}
		dfs(root)
		return
	}
	nums1, nums2 := inorder(root1), inorder(root2)
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			ans = append(ans, nums1[idx1])
			idx1++
		} else {
			ans = append(ans, nums2[idx2])
			idx2++
		}
	}
	for idx1 < len(nums1) {
		ans = append(ans, nums1[idx1])
		idx1++
	}
	for idx2 < len(nums2) {
		ans = append(ans, nums2[idx2])
		idx2++
	}
	return
}
