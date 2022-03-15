/*
-------------------------------------
# @Time    : 2021/11/10 16:57:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 347_前 K 个高频元素.go
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

func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	h := &IHeap347{}
	heap.Init(h)
	for key, value := range mp {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ans
}

type IHeap347 [][2]int

func (h IHeap347) Len() int {
	return len(h)
}

func (h IHeap347) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap347) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap347) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap347) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//func topKFrequent(nums []int, k int) (ans []int) {
//	mp := make(map[int]int)
//	for _, num := range nums {
//		mp[num]++
//	}
//	for key := range mp {
//		ans = append(ans, key)
//	}
//	// less function
//	sort.Slice(ans, func(x, y int) bool {
//		return mp[ans[x]] > mp[ans[y]]
//	})
//	return ans[:k]
//}
