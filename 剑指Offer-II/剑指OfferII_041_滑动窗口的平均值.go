/*
-------------------------------------
# @Time    : 2022/3/12 21:55:12
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_041_滑动窗口的平均值.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	ma := Constructor041(3)
	fmt.Println(ma.Next(1))
	fmt.Println(ma.Next(10))
	fmt.Println(ma.Next(3))
	fmt.Println(ma.Next(5))
}

type MovingAverage struct {
	queue []int
	sum   float64
	size  int
}

func Constructor041(size int) MovingAverage {
	return MovingAverage{[]int{}, 0.0, size}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) >= this.size {
		this.sum -= float64(this.queue[0])
		this.queue = this.queue[1:]
	}
	this.queue = append(this.queue, val)
	this.sum += float64(val)
	return this.sum / float64(len(this.queue))
}
