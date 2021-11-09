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

type MyStack struct {
	queue []int
}

func main() {
	stack := Constructor225()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.queue)
	fmt.Println(stack.Pop())
	fmt.Println(stack.queue)
	fmt.Println(stack.Empty())
}

func Constructor225() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	for i := 0; i < len(this.queue)-1; i++ {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.Push(val)
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.Push(val)
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
