/*
-------------------------------------
# @Time    : 2021/11/15 15:41:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 429_N 叉树的层序遍历.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &Node429{1, []*Node429{{3, []*Node429{{5, []*Node429{}}, {6, []*Node429{}}}}, {2, []*Node429{}}, {4, []*Node429{}}}}
	fmt.Println(levelOrderN(root))
}

func levelOrderN(root *Node429) (ans [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node429)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return
}

type Node429 struct {
	Val      int
	Children []*Node429
}
