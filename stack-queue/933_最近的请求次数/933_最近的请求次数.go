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
	rc := Constructor()
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

func Constructor() RecentCounter {
	return RecentCounter{0, 0, []int{}}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)
	rc.count++
	rc.last = t - 3000
	for rc.queue[0] < rc.last {
		rc.queue = rc.queue[1:]
		rc.count--
	}
	return rc.count
}
