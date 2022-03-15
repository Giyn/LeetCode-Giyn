/*
-------------------------------------
# @Time    : 2022/3/15 15:39:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_059_数据流的第 K 大数值.go
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
	kl := Constructor059(3, []int{4, 5, 8, 2})
	fmt.Println(kl.Add(3))
	fmt.Println(kl.Add(5))
	fmt.Println(kl.Add(10))
	fmt.Println(kl.Add(9))
	fmt.Println(kl.Add(4))
}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor059(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	a := this.IntSlice
	v := a[len(a)-1]
	this.IntSlice = a[:len(a)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
