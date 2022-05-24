/*
-------------------------------------
# @Time    : 2022/3/17 9:30:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 875_爱吃香蕉的珂珂.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{3, 6, 7, 11}
	H := 8
	fmt.Println(minEatingSpeed(piles, H))
}

func minEatingSpeed(piles []int, h int) int {
	canFinish := func(k int) bool {
		time := 0
		for _, pile := range piles {
			time += pile / k
			if pile%k > 0 {
				time += 1
			}
			if time > h {
				return false
			}
		}
		return true
	}
	left, right := 0, max(piles)
	for left <= right {
		mid := left + (right-left)>>1
		if mid == 0 {
			return 1
		}
		if canFinish(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func max(nums []int) int {
	ans := math.MinInt64
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans
}
