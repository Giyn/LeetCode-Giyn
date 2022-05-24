/*
-------------------------------------
# @Time    : 2021/12/20 9:10:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 475_供暖器.go
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
	houses := []int{1, 2, 3, 4}
	heaters := []int{1, 4}
	fmt.Println(findRadius(houses, heaters))
}

func findRadius(houses []int, heaters []int) (ans int) {
	sort.Ints(houses)
	sort.Ints(heaters)
	ans = math.MinInt64
	i := 0
	for _, house := range houses {
		for i < len(heaters) && heaters[i] < house {
			i++
		}
		if i == 0 {
			ans = max(ans, heaters[i]-house)
		} else if i == len(heaters) {
			return max(ans, houses[len(houses)-1]-heaters[len(heaters)-1])
		} else {
			ans = max(ans, min(heaters[i]-house, house-heaters[i-1]))
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
