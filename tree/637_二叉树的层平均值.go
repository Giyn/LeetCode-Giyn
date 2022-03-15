/*
-------------------------------------
# @Time    : 2021/11/15 15:22:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 637_二叉树的层平均值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(averageOfLevels(root))
}

func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []float64
		sum := 0.0
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, float64(node.Val))
		}
		for _, v := range tmp {
			sum += v
		}
		ans = append(ans, sum/float64(len(tmp)))
	}
	return
}
