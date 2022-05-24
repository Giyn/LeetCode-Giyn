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
	minStack := Constructor()
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

func Constructor() MinStack {
	return MinStack{minStack: []int{math.MaxInt64}, stack: []int{}}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if val < ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
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

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}
