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
	minStack := Constructor()
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

func Constructor() MinStack {
	return MinStack{[]int{math.MaxInt64}, []int{}}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if x < ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, x)
	} else {
		ms.minStack = append(ms.minStack, ms.minStack[len(ms.minStack)-1])
	}
}

func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.minStack = ms.minStack[:len(ms.minStack)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Min() int {
	return ms.minStack[len(ms.minStack)-1]
}
