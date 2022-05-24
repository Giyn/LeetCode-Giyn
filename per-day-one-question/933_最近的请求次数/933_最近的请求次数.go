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
	rc := Constructor()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))
}

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{[]int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)
	last := t - 3000
	for rc.queue[0] < last {
		rc.queue = rc.queue[1:]
	}
	return len(rc.queue)
}
