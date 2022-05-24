/*
-------------------------------------
# @Time    : 2022/3/15 16:43:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_061_和最小的 k 个数对.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) (ans [][]int) {
	h := &hp{nil, nums1, nums2}
	heap.Init(h)
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, struct{ i, j int }{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		top := heap.Pop(h).(struct{ i, j int })
		i, j := top.i, top.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < len(nums2) {
			heap.Push(h, struct{ i, j int }{i, j + 1})
		}
	}
	return
}

type hp struct {
	data         []struct{ i, j int }
	nums1, nums2 []int
}

func (h hp) Len() int {
	return len(h.data)
}

func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}

func (h hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(struct{ i, j int }))
}

func (h *hp) Pop() interface{} {
	tmp := h.data
	v := tmp[len(tmp)-1]
	h.data = tmp[:len(tmp)-1]
	return v
}
