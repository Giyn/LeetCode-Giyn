/*
-------------------------------------
# @Time    : 2022/3/15 16:01:28
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_060_出现频率最高的 k 个数字.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) (ans []int) {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	h := &hp060{}
	heap.Init(h)
	for key, value := range mp {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(h).([2]int)[0])
	}
	return
}

type hp060 [][2]int

func (h hp060) Len() int {
	return len(h)
}

func (h hp060) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h hp060) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp060) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *hp060) Pop() interface{} {
	tmp := *h
	v := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return v
}
