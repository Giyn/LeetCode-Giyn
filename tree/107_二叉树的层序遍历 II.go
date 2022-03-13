/*
-------------------------------------
# @Time    : 2021/11/15 14:46:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 107_二叉树的层序遍历 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode107{3, &TreeNode107{9, nil, nil}, &TreeNode107{20, &TreeNode107{15, nil, nil}, &TreeNode107{7, nil, nil}}}
	fmt.Println(levelOrderBottom(root))
}

func levelOrderBottom(root *TreeNode107) (ans [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode107)
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
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return
}

type TreeNode107 struct {
	Val   int
	Left  *TreeNode107
	Right *TreeNode107
}
