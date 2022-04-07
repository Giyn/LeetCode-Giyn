/*
-------------------------------------
# @Time    : 2022/4/7 17:00:31
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 692_前K个高频单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}

func topKFrequent(words []string, k int) []string {
	mp := make(map[string]int, len(words))
	for _, word := range words {
		mp[word]++
	}
	var ans []string
	for key := range mp {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(i, j int) bool {
		return mp[ans[i]] > mp[ans[j]] || mp[ans[i]] == mp[ans[j]] && ans[i] < ans[j]
	})
	return ans[:k]
}

//func topKFrequent(words []string, k int) []string {
//	mp := make(map[string]int, len(words))
//	for _, word := range words {
//		mp[word]++
//	}
//	h := &hp692{}
//	heap.Init(h)
//	for word, cnt := range mp {
//		heap.Push(h, wordFrequent{word, cnt})
//		if h.Len() > k {
//			heap.Pop(h)
//		}
//	}
//	ans := make([]string, k)
//	for i := k - 1; i >= 0; i-- {
//		ans[i] = heap.Pop(h).(wordFrequent).word
//	}
//	return ans
//}
//
//type wordFrequent struct {
//	word string
//	cnt  int
//}
//
//type hp692 []wordFrequent
//
//func (h hp692) Len() int {
//	return len(h)
//}
//
//func (h hp692) Less(i, j int) bool {
//	return h[i].cnt < h[j].cnt || h[i].cnt == h[j].cnt && h[i].word > h[j].word
//}
//
//func (h hp692) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *hp692) Push(x interface{}) {
//	*h = append(*h, x.(wordFrequent))
//}
//
//func (h *hp692) Pop() interface{} {
//	tmp := *h
//	x := tmp[len(tmp)-1]
//	*h = tmp[:len(tmp)-1]
//	return x
//}
