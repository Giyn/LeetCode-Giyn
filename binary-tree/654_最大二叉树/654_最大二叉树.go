/*
-------------------------------------
# @Time    : 2021/11/23 21:01:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 654_最大二叉树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
)

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	PrintBinaryTree(root)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	if len(nums) == 1 {
		root.Val = nums[0]
		return root
	}
	maxIdx := maxIndex(nums)
	root.Val = nums[maxIdx]
	root.Left = constructMaximumBinaryTree(nums[:maxIdx])
	root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return root
}

func maxIndex(nums []int) (maxIdx int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	return
}
