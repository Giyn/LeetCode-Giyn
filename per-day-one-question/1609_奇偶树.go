/*
-------------------------------------
# @Time    : 2021/12/25 1:32:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1609_奇偶树.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"container/list"
	"fmt"
	"math"
)

func main() {
	//root := &TreeNode{1, &TreeNode{10, &TreeNode{3, &TreeNode{12, nil, nil}, &TreeNode{8, nil, nil}}, nil}, &TreeNode{4, &TreeNode{7, &TreeNode{6, nil, nil}, nil}, &TreeNode{9, nil, &TreeNode{2, nil, nil}}}}
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}}}
	fmt.Println(isEvenOddTree(root))
}

func isEvenOddTree(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	var even = true
	for queue.Len() > 0 {
		length := queue.Len()
		var prev int
		if even {
			prev = math.MinInt64
		} else {
			prev = math.MaxInt64
		}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if even && (node.Val%2 != 1 || node.Val <= prev) {
				return false
			} else if !even && (node.Val%2 != 0 || node.Val >= prev) {
				return false
			}
			prev = node.Val
		}
		even = !even
	}
	return true
}
