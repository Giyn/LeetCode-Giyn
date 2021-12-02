/*
-------------------------------------
# @Time    : 2021/12/2 14:37:52
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 447_回旋镖的数量.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	points := [][]int{{0, 0}, {1, 0}, {2, 0}}
	fmt.Println(numberOfBoomerangs(points))
}

func numberOfBoomerangs(points [][]int) (ans int) {
	n := len(points)
	for _, i := range points {
		mp := make(map[int]int, n)
		for _, j := range points {
			distance2 := (i[0]-j[0])*(i[0]-j[0]) + (i[1]-j[1])*(i[1]-j[1])
			mp[distance2]++
			if mp[distance2] >= 2 {
				ans += 2 * (mp[distance2] - 1)
			}
		}
	}
	return
}
