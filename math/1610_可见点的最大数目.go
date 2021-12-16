/*
-------------------------------------
# @Time    : 2021/12/16 9:05:02
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1610_可见点的最大数目.go
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
	points := [][]int{{2, 1}, {2, 2}, {3, 4}, {1, 1}}
	angle := 90
	location := []int{1, 1}
	fmt.Println(visiblePoints(points, angle, location))
}

func visiblePoints(points [][]int, angle int, location []int) int {
	sameCnt := 0
	var polarDegrees []float64
	for _, p := range points {
		if p[0] == location[0] && p[1] == location[1] {
			sameCnt++
		} else {
			polarDegrees = append(polarDegrees, math.Atan2(float64(p[1]-location[1]), float64(p[0]-location[0])))
		}
	}
	sort.Float64s(polarDegrees)

	n := len(polarDegrees)
	for _, deg := range polarDegrees {
		polarDegrees = append(polarDegrees, deg+2*math.Pi)
	}

	maxCnt := 0
	right := 0
	degree := float64(angle) * math.Pi / 180
	for i, deg := range polarDegrees[:n] {
		for right < n*2 && polarDegrees[right] <= deg+degree {
			right++
		}
		if right-i > maxCnt {
			maxCnt = right - i
		}
	}
	return sameCnt + maxCnt
}
