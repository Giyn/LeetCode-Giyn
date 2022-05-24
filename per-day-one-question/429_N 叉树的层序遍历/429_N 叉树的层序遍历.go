/*
-------------------------------------
# @Time    : 2022/4/8 0:02:37
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
	root := &Node{Val: 1, Children: []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(levelOrder(root))
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var tmp []int
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return
}
