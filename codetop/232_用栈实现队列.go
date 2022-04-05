/*
-------------------------------------
# @Time    : 2021/11/9 16:14:00
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 232_用栈实现队列.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

type MyQueue struct {
	inStack  []int
	outStack []int
}

func main() {
	queue := Constructor232()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}

func Constructor232() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) in2out() {
	for len(this.inStack) != 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	val := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return val
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.outStack) == 0 && len(this.inStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
