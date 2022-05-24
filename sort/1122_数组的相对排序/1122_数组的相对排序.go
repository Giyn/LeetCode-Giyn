/*
-------------------------------------
# @Time    : 2022/3/17 0:48:16
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1122_数组的相对排序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	myRank := map[int]int{}
	for i, v := range arr2 {
		myRank[v] = i - len(arr2)
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		if r, ok := myRank[x]; ok {
			x = r
		}
		if r, ok := myRank[y]; ok {
			y = r
		}
		return x < y
	})
	return arr1
}
