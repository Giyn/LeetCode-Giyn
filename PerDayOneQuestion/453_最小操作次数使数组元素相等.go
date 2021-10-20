/*
-------------------------------------
# @Time    : 2021/10/20 10:46:18
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 453_最小操作次数使数组元素相等.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
}

func minMoves(nums []int) int {
	minV := math.MaxInt64
	sum := 0

	for _, v := range nums {
		if minV > v {
			minV = v
		}
		sum += v
	}

	return sum - len(nums)*minV
}
