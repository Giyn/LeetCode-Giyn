/*
-------------------------------------
# @Time    : 2022/3/15 15:03:43
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_058_日程表.go
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

func (mc *MyCalendarNode) insert(node *MyCalendarNode) bool {
	if node.start >= mc.end {
		if mc.right == nil {
			mc.right = node
			return true
		}
		return mc.right.insert(node)
	}
	if node.end <= mc.start {
		if mc.left == nil {
			mc.left = node
			return true
		}
		return mc.left.insert(node)
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
