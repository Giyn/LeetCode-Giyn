/*
-------------------------------------
# @Time    : 2022/3/9 9:27:30
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 798_得分最高的最小轮调.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 4, 0}
	fmt.Println(bestRotation(nums))
}

func bestRotation(nums []int) int {
	n := len(nums)
	diffs := make([]int, n)
	for i, num := range nums {
		low := (i + 1) % n
		high := (i - num + n + 1) % n
		diffs[low]++
		diffs[high]--
		if low >= high {
			diffs[0]++
		}
	}
	score, maxScore, idx := 0, 0, 0
	for i, diff := range diffs {
		score += diff
		if score > maxScore {
			maxScore, idx = score, i
		}
	}
	return idx
}
