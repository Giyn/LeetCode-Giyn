/*
-------------------------------------
# @Time    : 2021/11/27 14:42:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 108_将有序数组转换为二叉搜索树.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(root.Val)
}

func sortedArrayToBST(nums []int) *TreeNode108 {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode108{}
	if len(nums) == 1 {
		root.Val = nums[0]
		return root
	}
	middleIndex := len(nums) / 2
	root.Val = nums[middleIndex]
	root.Left = sortedArrayToBST(nums[:middleIndex])
	root.Right = sortedArrayToBST(nums[middleIndex+1:])
	return root
}

type TreeNode108 struct {
	Val   int
	Left  *TreeNode108
	Right *TreeNode108
}
