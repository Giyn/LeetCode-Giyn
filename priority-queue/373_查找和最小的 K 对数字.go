/*
-------------------------------------
# @Time    : 2022/1/14 19:06:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 373_查找和最小的 K 对数字.go
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

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp373{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair373{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair373)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair373{i, j + 1})
		}
	}
	return
}

type pair373 struct{ i, j int }
type hp373 struct {
	data         []pair373
	nums1, nums2 []int
}

func (h hp373) Len() int { return len(h.data) }
func (h hp373) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp373) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp373) Push(v interface{}) { h.data = append(h.data, v.(pair373)) }
func (h *hp373) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }
