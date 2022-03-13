/*
-------------------------------------
# @Time    : 2022/3/12 21:55:27
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_042_最近请求次数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	rc := Constructor042()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))
}

type RecentCounter struct {
	count int
	last  int
	queue []int
}

func Constructor042() RecentCounter {
	return RecentCounter{0, 0, []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	this.count++
	this.last = t - 3000
	for this.queue[0] < this.last {
		this.queue = this.queue[1:]
		this.count--
	}
	return this.count
}
