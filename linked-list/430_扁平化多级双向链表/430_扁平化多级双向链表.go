/*
-------------------------------------
# @Time    : 2022/3/11 19:24:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 430_扁平化多级双向链表.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	var root *Node
	root = &Node{Val: 1, Next: &Node{Val: 2, Prev: root}, Child: &Node{Val: 3}}
	res := flatten(root)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	cur := head
	for cur != nil {
		// 没有子节点
		if cur.Child == nil {
			cur = cur.Next
			continue
		}

		// 含有子节点，找到该子节点的末尾后继节点
		tmp := cur.Child
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		// 连接起来
		tmp.Next = cur.Next
		if cur.Next != nil {
			cur.Next.Prev = tmp
		}

		cur.Next, cur.Child.Prev = cur.Child, cur
		cur.Child = nil // 题目要求
	}
	return head
}
