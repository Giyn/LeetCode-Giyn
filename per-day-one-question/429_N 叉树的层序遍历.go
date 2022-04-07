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
	. "LeetCodeGiyn/utils/multiway-tree"
	"fmt"
)

func main() {
	root := &Node{Val: 1, Children: []*Node{{3, []*Node{{5, []*Node{}}, {6, []*Node{}}}}, {2, []*Node{}}, {4, []*Node{}}}}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		length := len(queue)
		var tmp []int
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			for _, child := range cur.Children {
				queue = append(queue, child)
			}
		}
		ans = append(ans, tmp)
	}
	return
}
