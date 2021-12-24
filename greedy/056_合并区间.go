/*
-------------------------------------
# @Time    : 2021/12/24 15:16:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 056_合并区间.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) (ans [][]int) {
	// 8ms
	//sort.Slice(intervals, func(i, j int) bool {
	//	if intervals[i][0] == intervals[j][0] {
	//		return intervals[i][1] < intervals[j][1]
	//	}
	//	return intervals[i][0] < intervals[j][0]
	//})
	//start, end := intervals[0][0], intervals[0][1]
	//for i := 1; i < len(intervals); i++ {
	//	if intervals[i][0] <= end && intervals[i][1] > end {
	//		end = intervals[i][1]
	//	} else if intervals[i][0] > end {
	//		ans = append(ans, []int{start, end})
	//		start = intervals[i][0]
	//		end = intervals[i][1]
	//	}
	//}
	//ans = append(ans, []int{start, end})
	//return

	// 28ms
	var max func(int, int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			intervals[i][1] = max(intervals[i][1], intervals[i+1][1])
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}
	return intervals
}
