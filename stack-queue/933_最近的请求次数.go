/*
-------------------------------------
# @Time    : 2022/3/13 10:02:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 933_最近的请求次数.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	rc := Constructor933()
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

func Constructor933() RecentCounter {
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
