/*
-------------------------------------
# @Time    : 2022/3/14 0:54:25
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 599_两个列表的最小索引总和.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}

func findRestaurant(list1 []string, list2 []string) (ans []string) {
	mp := make(map[string]int, len(list1))
	minSum := math.MaxInt64
	var idx []int
	for i, v1 := range list1 {
		mp[v1] = i
	}
	for i, v2 := range list2 {
		if v, ok := mp[v2]; ok && v+i == minSum {
			idx = append(idx, i)
		}
		if v, ok := mp[v2]; ok && v+i < minSum {
			minSum = mp[v2] + i
			idx = []int{i}
		}
	}
	for _, i := range idx {
		ans = append(ans, list2[i])
	}
	return
}
