/*
-------------------------------------
# @Time    : 2022/3/15 15:36:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 729_我的日程安排表 I.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	mc := Constructor()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(15, 25))
	fmt.Println(mc.Book(20, 30))
}

type MyCalendarNode struct {
	left, right *MyCalendarNode
	start, end  int
}

func (n *MyCalendarNode) insert(node *MyCalendarNode) bool {
	if node.start >= n.end {
		if n.right == nil {
			n.right = node
			return true
		}
		return n.right.insert(node)
	}
	if node.end <= n.start {
		if n.left == nil {
			n.left = node
			return true
		}
		return n.left.insert(node)
	}
	return false
}

type MyCalendar struct {
	root *MyCalendarNode
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	node := &MyCalendarNode{start: start, end: end}
	if mc.root == nil {
		mc.root = node
		return true
	}
	return mc.root.insert(node)
}
