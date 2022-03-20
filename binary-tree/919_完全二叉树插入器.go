/*
-------------------------------------
# @Time    : 2022/3/13 10:42:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 919_完全二叉树插入器.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/binary-tree"
	"fmt"
)

func main() {
	cbt := Constructor919(&TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}})
	fmt.Println(cbt.Insert(7))
	fmt.Println(cbt.Insert(8))
	fmt.Println(cbt.Get_root())
}

type CBTInserter struct {
	queue []*TreeNode
}

func Constructor919(root *TreeNode) CBTInserter {
	this := CBTInserter{[]*TreeNode{root}}
	for i := 0; i < len(this.queue); i++ {
		if this.queue[i].Left != nil {
			this.queue = append(this.queue, this.queue[i].Left)
		}
		if this.queue[i].Right != nil {
			this.queue = append(this.queue, this.queue[i].Right)
		}
	}
	return this
}

func (this *CBTInserter) Insert(v int) int {
	n := len(this.queue)
	father := this.queue[(n-1)>>1]
	this.queue = append(this.queue, &TreeNode{Val: v})
	if n&1 == 1 {
		father.Left = this.queue[n]
	} else {
		father.Right = this.queue[n]
	}
	return father.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.queue[0]
}
