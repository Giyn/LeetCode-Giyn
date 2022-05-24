/*
-------------------------------------
# @Time    : 2021/11/9 21:40:17
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 225_用队列实现栈.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.queue)
	fmt.Println(stack.Pop())
	fmt.Println(stack.queue)
	fmt.Println(stack.Empty())
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (ms *MyStack) Push(x int) {
	ms.queue = append(ms.queue, x)
}

func (ms *MyStack) Pop() int {
	for i := 0; i < len(ms.queue)-1; i++ {
		val := ms.queue[0]
		ms.queue = ms.queue[1:]
		ms.Push(val)
	}
	val := ms.queue[0]
	ms.queue = ms.queue[1:]
	return val
}

func (ms *MyStack) Top() int {
	val := ms.Pop()
	ms.Push(val)
	return val
}

func (ms *MyStack) Empty() bool {
	return len(ms.queue) == 0
}
