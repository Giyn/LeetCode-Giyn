/*
-------------------------------------
# @Time    : 2022/4/30 0:15:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 908_最小差值 I.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 6}
	k := 3
	fmt.Println(smallestRangeI(nums, k))
}

func smallestRangeI(nums []int, k int) int {
	max, min := math.MinInt64, math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	if max-min <= 2*k {
		return 0
	}
	return max - min - 2*k
}
