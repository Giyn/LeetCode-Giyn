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
	q := Constructor009()
	q.AppendTail(3)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}

type CQueue struct {
	outStack []int
	inStack  []int
}

func Constructor009() CQueue {
	return CQueue{outStack: []int{}, inStack: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	if len(this.outStack) == 0 {
		return -1
	}
	v := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return v
}
