/*
-------------------------------------
# @Time    : 2022/3/11 19:26:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 155_最小栈.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	minStack := Constructor155()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor155() MinStack {
	return MinStack{minStack: []int{math.MaxInt64}, stack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[len(this.minStack)-1])
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
