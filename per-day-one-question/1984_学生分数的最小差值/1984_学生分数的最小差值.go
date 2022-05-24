/*
-------------------------------------
# @Time    : 2022/2/11 0:09:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1984_学生分数的最小差值.go
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
	nums := []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
	k := 6
	fmt.Println(minimumDifference(nums, k))
}

func minimumDifference(nums []int, k int) (ans int) {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans = math.MaxInt64
	sort.Ints(nums)
	for left, right := 0, k-1; right < len(nums); left, right = left+1, right+1 {
		ans = min(ans, nums[right]-nums[left])
	}
	return
}
