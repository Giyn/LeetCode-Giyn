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
	. "LeetCodeGiyn/utils/multiway-tree"
	"container/list"
	"fmt"
)

func main() {
	root := &Node{1, []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(levelOrderN(root))
}

func levelOrderN(root *Node) (ans [][]int) {
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
