/*
-------------------------------------
# @Time    : 2021/11/29 11:43:32
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 215_数组中的第K个最大元素.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	h := new(hp)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for _, num := range nums[k:] {
		top := heap.Pop(h).(int)
		if num > top {
			heap.Push(h, num)
		} else {
			heap.Push(h, top)
		}
	}
	return heap.Pop(h).(int)
}

type hp []int

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *hp) Pop() interface{}   { t := *h; x := t[len(t)-1]; *h = t[:len(t)-1]; return x }
