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
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(root.Val)
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
	maxIndex := max654(nums)
	root.Val = nums[maxIndex]
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

func max654(nums []int) (maxIndex int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return
}
