/*
-------------------------------------
# @Time    : 2021/12/16 23:59:47
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1005_K 次取反后最大化的数组和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	fmt.Println(largestSumAfterKNegations(nums, k))
}

func largestSumAfterKNegations(nums []int, k int) (ans int) {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	fmt.Println(nums)
	for _, num := range nums {
		ans += num
	}
	return
}
