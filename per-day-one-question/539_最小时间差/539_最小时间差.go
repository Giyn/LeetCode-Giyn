/*
-------------------------------------
# @Time    : 2022/1/18 21:44:24
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 539_最小时间差.go
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
	timePoints := []string{"00:00", "23:59", "00:00"}
	fmt.Println(findMinDifference(timePoints))
}

func findMinDifference(timePoints []string) (ans int) {
	var getMinutes func(string) int
	getMinutes = func(s string) int {
		return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'0')*10 + int(s[4]-'0')
	}
	if len(timePoints) > 1440 {
		return 0
	}
	sort.Strings(timePoints)
	ans = math.MaxInt64
	first := getMinutes(timePoints[0])
	pre := first
	for _, t := range timePoints[1:] {
		minute := getMinutes(t)
		ans = min(ans, minute-pre)
		pre = minute
	}
	ans = min(ans, first+1440-pre)
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
