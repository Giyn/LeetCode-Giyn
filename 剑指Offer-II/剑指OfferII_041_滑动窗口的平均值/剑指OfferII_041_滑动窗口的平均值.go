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
	ma := Constructor(3)
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

func Constructor(size int) MovingAverage {
	return MovingAverage{[]int{}, 0.0, size}
}

func (mv *MovingAverage) Next(val int) float64 {
	if len(mv.queue) >= mv.size {
		mv.sum -= float64(mv.queue[0])
		mv.queue = mv.queue[1:]
	}
	mv.queue = append(mv.queue, val)
	mv.sum += float64(val)
	return mv.sum / float64(len(mv.queue))
}
