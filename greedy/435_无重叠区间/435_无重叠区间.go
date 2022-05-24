/*
-------------------------------------
# @Time    : 2021/12/21 22:25:06
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 435_无重叠区间.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	fmt.Println(intervals)
	count := 1
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}
	return len(intervals) - count
}
