/*
-------------------------------------
# @Time    : 2022/3/12 19:18:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指OfferII_035_最小时间差.go
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
	getMinutes := func(s string) int {
		return (int(s[0]-'0')*10+int(s[1]-'0'))*60 + int(s[3]-'a')*10 + int(s[4]-'a')
	}
	if len(timePoints) > 1440 {
		return 0
	}
	sort.Strings(timePoints)
	first := getMinutes(timePoints[0])
	pre := first
	ans = math.MaxInt64
	for _, timePoint := range timePoints[1:] {
		minutes := getMinutes(timePoint)
		ans = min(ans, minutes-pre)
		pre = minutes
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
