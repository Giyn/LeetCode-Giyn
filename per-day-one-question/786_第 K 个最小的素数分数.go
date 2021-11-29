/*
-------------------------------------
# @Time    : 2021/11/29 0:58:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 786_第 K 个最小的素数分数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	fmt.Println(kthSmallestPrimeFraction(arr, k))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp786, n-1)
	for j := 1; j < n; j++ {
		h[j-1] = frac{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)
	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(frac)
		if f.i+1 < f.j {
			heap.Push(&h, frac{arr[f.i+1], f.y, f.i + 1, f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type hp786 []frac

func (h hp786) Len() int            { return len(h) }
func (h hp786) Less(i, j int) bool  { return h[i].x*h[j].y < h[i].y*h[j].x } // 防止精度丢失 交叉相乘
func (h hp786) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp786) Push(v interface{}) { *h = append(*h, v.(frac)) }
func (h *hp786) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
