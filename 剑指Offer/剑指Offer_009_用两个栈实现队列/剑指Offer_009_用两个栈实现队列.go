/*
-------------------------------------
# @Time    : 2022/3/25 12:20:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_009_用两个栈实现队列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	q := Constructor()
	q.AppendTail(3)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}

type CQueue struct {
	outStack []int
	inStack  []int
}

func Constructor() CQueue {
	return CQueue{outStack: []int{}, inStack: []int{}}
}

func (cq *CQueue) AppendTail(value int) {
	cq.inStack = append(cq.inStack, value)
}

func (cq *CQueue) DeleteHead() int {
	if len(cq.outStack) == 0 {
		for len(cq.inStack) > 0 {
			cq.outStack = append(cq.outStack, cq.inStack[len(cq.inStack)-1])
			cq.inStack = cq.inStack[:len(cq.inStack)-1]
		}
	}
	if len(cq.outStack) == 0 {
		return -1
	}
	v := cq.outStack[len(cq.outStack)-1]
	cq.outStack = cq.outStack[:len(cq.outStack)-1]
	return v
}
