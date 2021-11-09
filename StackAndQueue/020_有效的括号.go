/*
-------------------------------------
# @Time    : 2021/11/9 22:52:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 020_有效的括号.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "([)]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := Constructor020()
	for i := 0; i < len(s); i++ {
		if (!stack.Empty() && stack.Top() == '(' && s[i] == ')') || (!stack.Empty() && stack.Top() == '{' && s[i] == '}') || (!stack.Empty() && stack.Top() == '[' && s[i] == ']') {
			stack.Pop()
			continue
		}
		stack.Push(s[i])
	}
	return stack.Empty()
}

type Stack020 struct {
	queue []byte
}

func Constructor020() Stack020 {
	return Stack020{
		queue: make([]byte, 0),
	}
}

func (this *Stack020) Push(x byte) {
	this.queue = append(this.queue, x)
}

func (this *Stack020) Pop() byte {
	for i := 0; i < len(this.queue)-1; i++ {
		val := this.queue[0]
		this.queue = this.queue[1:]
		this.Push(val)
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *Stack020) Top() byte {
	val := this.Pop()
	this.Push(val)
	return val
}

func (this *Stack020) Empty() bool {
	return len(this.queue) == 0
}
