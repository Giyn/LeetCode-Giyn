/*
-------------------------------------
# @Time    : 2022/1/12 23:21:01
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 334_递增的三元子序列.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums))
}

func increasingTriplet(nums []int) bool {
	var a, b = math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if num <= a {
			a = num
		} else if num <= b {
			b = num
		} else {
			return true
		}
	}
	return false
}
