/*
-------------------------------------
# @Time    : 2022/3/17 9:31:29
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_076_数组中的第 k 大的数字.go
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
	h := &hp076{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for _, num := range nums[k:] {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}

type hp076 struct {
	nums []int
}

func (h hp076) Len() int {
	return len(h.nums)
}

func (h hp076) Less(i, j int) bool {
	return h.nums[i] < h.nums[j]
}

func (h hp076) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *hp076) Push(x interface{}) {
	h.nums = append(h.nums, x.(int))
}

func (h *hp076) Pop() interface{} {
	tmp := h.nums
	v := tmp[len(tmp)-1]
	h.nums = tmp[:len(tmp)-1]
	return v
}
