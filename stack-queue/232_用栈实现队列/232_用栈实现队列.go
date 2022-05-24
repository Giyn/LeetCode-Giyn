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

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Peek())
	fmt.Println(queue.Empty())
}

type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (mq *MyQueue) Push(x int) {
	mq.inStack = append(mq.inStack, x)
}

func (mq *MyQueue) in2out() {
	for len(mq.inStack) != 0 {
		mq.outStack = append(mq.outStack, mq.inStack[len(mq.inStack)-1])
		mq.inStack = mq.inStack[:len(mq.inStack)-1]
	}
}

func (mq *MyQueue) Pop() int {
	if len(mq.outStack) == 0 {
		mq.in2out()
	}
	val := mq.outStack[len(mq.outStack)-1]
	mq.outStack = mq.outStack[:len(mq.outStack)-1]
	return val
}

func (mq *MyQueue) Peek() int {
	if len(mq.outStack) == 0 {
		mq.in2out()
	}
	return mq.outStack[len(mq.outStack)-1]
}

func (mq *MyQueue) Empty() bool {
	if len(mq.outStack) == 0 && len(mq.inStack) == 0 {
		return true
	}
	return false
}
