/*
-------------------------------------
# @Time    : 2022/2/7 19:51:34
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1405_最长快乐字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	a, b, c := 1, 1, 7
	fmt.Println(longestDiverseString(a, b, c))
}

type CharNum struct {
	val  int
	char byte
}

func longestDiverseString(a int, b int, c int) string {
	pq, ans := &IntHeap{}, []byte{}
	heap.Push(pq, CharNum{a, 'a'})
	heap.Push(pq, CharNum{b, 'b'})
	heap.Push(pq, CharNum{c, 'c'})
	for pq.Len() > 0 {
		v := addChar(pq, ans)
		if len(ans) == len(v) {
			break
		}
		ans = v
	}
	return string(ans)
}

func addChar(pq *IntHeap, ans []byte) []byte {
	cur := heap.Pop(pq).(CharNum)
	if cur.val == 0 {
		return ans
	}
	if len(ans) >= 2 && cur.char == ans[len(ans)-1] && cur.char == ans[len(ans)-2] {
		if pq.Len() == 0 {
			return ans
		}
		v := addChar(pq, ans)
		if len(ans) == len(v) {
			return ans
		}
		ans = v
	}
	ans = append(ans, cur.char)
	cur.val--
	heap.Push(pq, cur)
	return ans
}

type IntHeap []CharNum

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(CharNum))
}
