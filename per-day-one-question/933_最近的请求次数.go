/*
-------------------------------------
# @Time    : 2022/5/6 0:15:32
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
	queue []int
}

func Constructor933() RecentCounter {
	return RecentCounter{[]int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	last := t - 3000
	for this.queue[0] < last {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}
