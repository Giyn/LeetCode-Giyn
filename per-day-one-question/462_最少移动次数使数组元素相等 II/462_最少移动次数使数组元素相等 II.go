/*
-------------------------------------
# @Time    : 2022/5/19 0:19:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 462_最少移动次数使数组元素相等 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}

func minMoves(nums []int) (ans int) {
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	for _, num := range nums {
		ans += abs(num - mid)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
