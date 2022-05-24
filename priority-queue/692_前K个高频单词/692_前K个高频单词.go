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
	"container/heap"
	"fmt"
)

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequentWords(words, k))
}

func topKFrequentWords(words []string, k int) []string {
	mp := make(map[string]int, len(words))
	for _, word := range words {
		mp[word]++
	}
	h := &hp{}
	heap.Init(h)
	for word, cnt := range mp {
		heap.Push(h, wordFrequent{word, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(wordFrequent).word
	}
	return ans
}

type wordFrequent struct {
	word string
	cnt  int
}

type hp []wordFrequent

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt || h[i].cnt == h[j].cnt && h[i].word > h[j].word
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(wordFrequent))
}

func (h *hp) Pop() interface{} {
	tmp := *h
	x := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return x
}
