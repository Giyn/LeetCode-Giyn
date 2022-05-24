/*
-------------------------------------
# @Time    : 2022/3/13 10:03:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_043_往完全二叉树添加节点.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	cbt := Constructor(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}})
	fmt.Println(cbt.Insert(7))
	fmt.Println(cbt.Insert(8))
	fmt.Println(cbt.Get_root())
}

type CBTInserter struct {
	queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	cbt := CBTInserter{[]*TreeNode{root}}
	for i := 0; i < len(cbt.queue); i++ {
		if cbt.queue[i].Left != nil {
			cbt.queue = append(cbt.queue, cbt.queue[i].Left)
		}
		if cbt.queue[i].Right != nil {
			cbt.queue = append(cbt.queue, cbt.queue[i].Right)
		}
	}
	return cbt
}

func (cbt *CBTInserter) Insert(v int) int {
	n := len(cbt.queue)
	father := cbt.queue[(n-1)>>1]
	cbt.queue = append(cbt.queue, &TreeNode{Val: v})
	if n&1 == 1 {
		father.Left = cbt.queue[n]
	} else {
		father.Right = cbt.queue[n]
	}
	return father.Val
}

func (cbt *CBTInserter) Get_root() *TreeNode {
	return cbt.queue[0]
}
