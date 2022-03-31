/*
-------------------------------------
# @Time    : 2022/3/31 22:54:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_030_包含min函数的栈.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	minStack := Constructor030()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.Min())
}

type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor030() MinStack {
	return MinStack{[]int{math.MaxInt64}, []int{}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if x < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
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

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}
