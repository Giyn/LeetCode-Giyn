/*
-------------------------------------
# @Time    : 2022/9/21 23:56:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 面试题59_II_队列的最大值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	mq := Constructor()
	mq.Push_back(1)
	mq.Push_back(2)
	fmt.Println(mq.Max_value())
	fmt.Println(mq.Pop_front())
	fmt.Println(mq.Max_value())
}

type MaxQueue struct {
	q []int
	p []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.q) == 0 {
		return -1
	}
	return this.p[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q = append(this.q, value)
	for len(this.p) > 0 && value > this.p[len(this.p)-1] {
		this.p = this.p[:len(this.p)-1]
	}
	this.p = append(this.p, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.q) == 0 {
		return -1
	}
	if this.q[0] == this.p[0] {
		this.p = this.p[1:]
	}
	value := this.q[0]
	this.q = this.q[1:]
	return value
}
