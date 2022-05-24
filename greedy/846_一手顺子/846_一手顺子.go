/*
-------------------------------------
# @Time    : 2021/12/30 9:43:14
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 846_一手顺子.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	hand := []int{1, 2, 3, 4, 5}
	groupSize := 4
	fmt.Println(isNStraightHand(hand, groupSize))
}

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	mp := make(map[int]int)
	for _, v := range hand {
		mp[v]++
	}
	for _, x := range hand {
		if v, ok := mp[x]; ok && v > 0 {
			for i := x; i < x+groupSize; i++ {
				if mp[i] == 0 {
					return false
				}
				mp[i]--
			}
		}
	}
	return true
}
