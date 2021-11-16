/*
-------------------------------------
# @Time    : 2021/11/15 23:38:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 104_二叉树的最大深度.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode104{3, &TreeNode104{9, nil, nil}, &TreeNode104{20, &TreeNode104{15, nil, nil}, &TreeNode104{7, nil, nil}}}
	fmt.Println(maxDepth(root))
}

func maxDepth(root *TreeNode104) (ans int) {
	if root == nil {
		return 0
	}
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode104)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		depth++
	}
	return depth
}

type TreeNode104 struct {
	Val   int
	Left  *TreeNode104
	Right *TreeNode104
}
