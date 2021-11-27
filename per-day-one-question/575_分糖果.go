/*
-------------------------------------
# @Time    : 2021/11/1 0:02:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 575_分糖果.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	//candyType := []int{1, 1, 2, 2, 3, 3}
	candyType := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candyType))
}

func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	count := 1
	for i := 1; i < len(candyType) && count < len(candyType)/2; i++ {
		if candyType[i] > candyType[i-1] {
			count++
		}
	}
	return count
}

//func distributeCandies(candyType []int) int {
//	mp := make(map[int]int)
//
//	for _, candy := range candyType {
//		mp[candy]++
//	}
//
//	if len(mp) >= len(candyType)/2 {
//		return len(candyType) / 2
//	} else {
//		return len(mp)
//	}
//}
