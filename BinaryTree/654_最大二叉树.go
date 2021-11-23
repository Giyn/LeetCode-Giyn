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
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	root := constructMaximumBinaryTree(nums)
	fmt.Println(root.Val)
}

func constructMaximumBinaryTree(nums []int) *TreeNode654 {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode654{}
	if len(nums) == 1 {
		root.Val = nums[0]
		return root
	}
	tmp := max654(nums)
	maxIndex := tmp[0]
	root.Val = tmp[1]
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

type TreeNode654 struct {
	Val   int
	Left  *TreeNode654
	Right *TreeNode654
}

func max654(nums []int) [2]int {
	max := math.MinInt64
	maxIndex := -1
	for i, num := range nums {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return [2]int{maxIndex, max}
}
