/*
-------------------------------------
# @Time    : 2021/11/25 14:18:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 501_二叉搜索树中的众数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{2, nil, nil}, nil}}
	fmt.Println(findMode(root))
}

func findMode(root *TreeNode) (ans []int) {
	// 递归
	curCount := 0
	maxCount := 0
	var pre *TreeNode
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		if pre != nil && pre.Val == node.Val {
			curCount++
		} else {
			curCount = 1
		}
		if curCount > maxCount {
			maxCount = curCount
			ans = []int{node.Val}
		} else if curCount == maxCount {
			ans = append(ans, node.Val)
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return

	// 迭代
	//stack := list.New()
	//cur := root
	//curVal := 0
	//curCount := 0
	//maxCount := 0
	//for cur != nil || stack.Len() != 0 {
	//	if cur != nil {
	//		stack.PushBack(cur)
	//		cur = cur.Left
	//	} else {
	//		cur = stack.Remove(stack.Back()).(*TreeNode)
	//		// 判断是否有新的数字
	//		if cur.Val == curVal {
	//			curCount++
	//		} else {
	//			curCount = 1
	//			curVal = cur.Val
	//		}
	//		// 更新众数
	//		if curCount > maxCount {
	//			maxCount = curCount
	//			ans = []int{curVal}
	//		} else if curCount == maxCount {
	//			ans = append(ans, curVal)
	//		}
	//		cur = cur.Right
	//	}
	//}
	//return
}
