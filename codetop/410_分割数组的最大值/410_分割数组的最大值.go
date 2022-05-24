/*
-------------------------------------
# @Time    : 2022/4/6 0:51:19
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 410_分割数组的最大值.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(splitArray(nums, m))
}

func splitArray(nums []int, m int) int {
	// 答案在 [max(nums), sum(nums)] 之间
	left, right := max(nums), 0
	for _, num := range nums {
		right += num
	}
	for left <= right {
		mid := left + (right-left)>>1
		cnt, sum := 0, 0
		// 开始分组
		for _, num := range nums {
			sum += num
			if sum > mid {
				cnt++
				sum = num
			}
		}
		cnt++ // 最后一组
		if cnt > m {
			left = mid + 1
		} else {
			right = mid - 1
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
