/*
-------------------------------------
# @Time    : 2021/11/15 16:01:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 515_在每个树行中找最大值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"container/list"
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}}
	fmt.Println(largestValues(root))
}

func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		max := math.MinInt64
		for _, v := range tmp {
			if v > max {
				max = v
			}
		}
		ans = append(ans, max)
	}
	return
}
