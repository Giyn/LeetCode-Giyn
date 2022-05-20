/*
-------------------------------------
# @Time    : 2022/5/20 18:26:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_041_数据流中的中位数.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	mf := Constructor041()
	mf.AddNum(1)
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
}

type MedianFinder struct {
	queMin, queMax hp041
}

func Constructor041() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	minQ, maxQ := &this.queMin, &this.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	minQ, maxQ := this.queMin, this.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}
	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp041 struct {
	sort.IntSlice
}

func (hp *hp041) Pop() interface{} {
	tmp := hp.IntSlice
	x := tmp[len(tmp)-1]
	hp.IntSlice = tmp[:len(tmp)-1]
	return x
}

func (hp *hp041) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}
