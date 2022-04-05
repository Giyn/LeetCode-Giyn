/*
-------------------------------------
# @Time    : 2022/4/5 22:01:09
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 628_三个数的最大乘积.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(maximumProduct(nums))
}

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, x := range nums {
		if x < min1 {
			min2 = min1
			min1 = x
		} else if x < min2 {
			min2 = x
		}
		if x > max1 {
			max3 = max2
			max2 = max1
			max1 = x
		} else if x > max2 {
			max3 = max2
			max2 = x
		} else if x > max3 {
			max3 = x
		}
	}
	return Max(min1*min2*max1, max1*max2*max3)
}
