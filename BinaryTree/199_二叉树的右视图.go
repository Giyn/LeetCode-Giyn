/*
-------------------------------------
# @Time    : 2021/11/15 15:10:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 199_二叉树的右视图.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode199{1, &TreeNode199{2, nil, &TreeNode199{5, nil, nil}}, &TreeNode199{3, nil, &TreeNode199{4, nil, nil}}}
	fmt.Println(rightSideView(root))
}

func rightSideView(root *TreeNode199) (ans []int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode199)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp[len(tmp)-1])
	}
	return
}

type TreeNode199 struct {
	Val   int
	Left  *TreeNode199
	Right *TreeNode199
}
