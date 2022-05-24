/*
-------------------------------------
# @Time    : 2022/5/17 4:46:49
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_040_最小的k个数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{3, 2, 1}
	k := 2
	fmt.Println(getLeastNumbers(arr, k))
}

func getLeastNumbers(arr []int, k int) []int {
	if k <= 0 || len(arr) < k {
		return []int{}
	}
	h := &hp{}
	for _, num := range arr {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			top := heap.Pop(h).(int)
			if top > num {
				heap.Push(h, num)
			} else {
				heap.Push(h, top)
			}
		}
	}
	return *h
}

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	tmp := *h
	x := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return x
}
