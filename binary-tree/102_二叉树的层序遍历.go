/*
-------------------------------------
# @Time    : 2021/11/15 14:33:54
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 102_二叉树的层序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode102{3, &TreeNode102{9, nil, nil}, &TreeNode102{20, &TreeNode102{15, nil, nil}, &TreeNode102{7, nil, nil}}}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode102) (ans [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode102)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return
}

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}
