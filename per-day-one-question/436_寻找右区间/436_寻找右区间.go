/*
-------------------------------------
# @Time    : 2022/5/20 2:28:55
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 436_寻找右区间.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 4}, {2, 3}, {3, 4}}
	fmt.Println(findRightInterval(intervals))
}

func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	n := len(intervals)
	ans := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool { return intervals[i][0] >= p[1] })
		if i < n {
			ans[p[2]] = intervals[i][2]
		} else {
			ans[p[2]] = -1
		}
	}
	return ans
}
