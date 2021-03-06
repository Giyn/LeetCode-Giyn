/*
-------------------------------------
# @Time    : 2021/12/24 9:44:51
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1705_吃苹果的最大数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	apples := []int{1, 2, 3, 5, 2}
	days := []int{3, 2, 1, 4, 2}
	fmt.Println(eatenApples(apples, days))
}

func eatenApples(apples []int, days []int) (ans int) {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	hp := hp{}
	i := 0
	for ; i < len(apples); i++ {
		for hp.Len() > 0 && hp[0].end <= i {
			heap.Pop(&hp)
		}
		if apples[i] > 0 {
			heap.Push(&hp, pair{i + days[i], apples[i]})
		}
		if hp.Len() > 0 {
			hp[0].left--
			if hp[0].left == 0 {
				heap.Pop(&hp)
			}
			ans++
		}
	}
	for hp.Len() > 0 {
		for hp.Len() > 0 && hp[0].end <= i {
			heap.Pop(&hp)
		}
		if hp.Len() == 0 {
			break
		}
		tmp := heap.Pop(&hp).(pair)
		nums := min(tmp.end-i, tmp.left)
		ans += nums
		i += nums
	}
	return
}

type pair struct {
	end, left int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].end < h[j].end
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	tmp := *h
	x := tmp[len(tmp)-1]
	*h = tmp[:len(tmp)-1]
	return x
}
