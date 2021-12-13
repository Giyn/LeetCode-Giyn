/*
-------------------------------------
# @Time    : 2021/12/14 0:55:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 630_课程表 III.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	courses := [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}
	fmt.Println(scheduleCourse(courses))
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	h := &courseHeap{}
	total := 0 // 优先队列中的课程总时间
	for _, course := range courses {
		if t := course[0]; total+t <= course[1] {
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] {
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type courseHeap struct {
	sort.IntSlice
}

func (c courseHeap) Less(i, j int) bool {
	return c.IntSlice[i] > c.IntSlice[j]
}

func (c *courseHeap) Push(x interface{}) {
	c.IntSlice = append(c.IntSlice, x.(int))
}

func (c *courseHeap) Pop() interface{} {
	a := c.IntSlice
	x := a[len(a)-1]
	c.IntSlice = a[:len(a)-1]
	return x
}
